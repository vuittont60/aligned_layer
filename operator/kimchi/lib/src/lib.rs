use std::io::ErrorKind;
use std::{io::Error, mem, sync::Arc};

use kimchi::groupmap::GroupMap;
use kimchi::mina_curves::pasta::{Fp, VestaParameters};
use kimchi::mina_poseidon::constants::PlonkSpongeConstantsKimchi;
use kimchi::mina_poseidon::sponge::{DefaultFqSponge, DefaultFrSponge};
use kimchi::poly_commitment::commitment::CommitmentCurve;
use kimchi::verifier::verify;
use kimchi::{
    curve::KimchiCurve,
    mina_curves::pasta::Vesta,
    poly_commitment::{evaluation_proof::OpeningProof, srs::SRS},
    proof::ProverProof,
    verifier_index::VerifierIndex,
};

const MAX_PROOF_SIZE: usize = 10 * 1024;
const MAX_PUB_INPUT_SIZE: usize = 3 * 1024 * 1024;

type SpongeParams = PlonkSpongeConstantsKimchi;
type BaseSponge = DefaultFqSponge<VestaParameters, SpongeParams>;
type ScalarSponge = DefaultFrSponge<Fp, SpongeParams>;

#[no_mangle]
pub extern "C" fn verify_kimchi_proof_ffi(
    proof_bytes: &[u8; MAX_PROOF_SIZE],
    proof_len: usize,
    pub_input_bytes: &[u8; MAX_PUB_INPUT_SIZE],
    pub_input_len: usize,
) -> bool {
    let proof = if let Ok(proof) = rmp_serde::from_slice(&proof_bytes[..proof_len]) {
        proof
    } else {
        return false;
    };

    let verifier_index = if let Ok(verifier_index) =
        deserialize_kimchi_pub_input(pub_input_bytes[..pub_input_len].to_vec())
    {
        verifier_index
    } else {
        return false;
    };

    let group_map = <Vesta as CommitmentCurve>::Map::setup();

    verify::<Vesta, BaseSponge, ScalarSponge, OpeningProof<Vesta>>(
        &group_map,
        &verifier_index,
        &proof,
        &Vec::new(),
    )
    .is_ok()
}

pub fn serialize_kimchi_pub_input(
    verifier_index: &VerifierIndex<Vesta, OpeningProof<Vesta>>,
    srs: &SRS<Vesta>,
) -> Vec<u8> {
    let mut pub_input: Vec<u8> = Vec::new();
    let verifier_index_bytes =
        bincode::serialize(verifier_index).expect("Could not serialize verifier index");
    let verifier_index_bytes_len = verifier_index_bytes.len() as u32;
    let srs_bytes = bincode::serialize(srs).expect("Could not serialize SRS");
    let srs_bytes_len = srs_bytes.len() as u32;

    pub_input.extend_from_slice(&verifier_index_bytes_len.to_be_bytes());
    pub_input.extend_from_slice(&verifier_index_bytes);
    pub_input.extend_from_slice(&srs_bytes_len.to_be_bytes());
    pub_input.extend_from_slice(&srs_bytes);

    pub_input
}

fn deserialize_kimchi_pub_input(
    pub_input_bytes: Vec<u8>,
) -> Result<VerifierIndex<Vesta, OpeningProof<Vesta>>, Box<dyn std::error::Error>> {
    let mut pub_input_bytes = pub_input_bytes;
    let u32_bytes_size = mem::size_of::<u32>();

    // verifier index deserialization
    let verifier_index_bytes_len_bytes: Vec<u8> = pub_input_bytes.drain(..u32_bytes_size).collect();
    let verifier_index_bytes_len =
        u32::from_be_bytes(verifier_index_bytes_len_bytes.as_slice().try_into()?) as usize;
    let verifier_index_bytes: Vec<u8> = pub_input_bytes.drain(..verifier_index_bytes_len).collect();
    let mut verifier_index: VerifierIndex<Vesta, OpeningProof<Vesta>> =
        bincode::deserialize(&verifier_index_bytes)?;

    // srs deserialization
    let srs_bytes_len_bytes: Vec<u8> = pub_input_bytes.drain(..u32_bytes_size).collect();
    let srs_bytes_len = u32::from_be_bytes(srs_bytes_len_bytes.as_slice().try_into()?) as usize;
    let srs_bytes: Vec<u8> = pub_input_bytes.drain(..srs_bytes_len).collect();
    let mut srs: SRS<Vesta> = bincode::deserialize(&srs_bytes)?;

    if !(pub_input_bytes.is_empty()) {
        return Err(Box::new(Error::new(
            ErrorKind::Other,
            "Public input buffer should be empty",
        )));
    }

    // add necessary fields to verifier index
    srs.add_lagrange_basis(verifier_index.domain);

    // we only need srs to be embedded in the verifier index, so no need to return it
    verifier_index.srs = Arc::new(srs);
    verifier_index.endo = *Vesta::other_curve_endo();

    Ok(verifier_index)
}

#[cfg(test)]
mod test {
    use std::{io::BufReader, path::Path};

    use kimchi::groupmap::GroupMap;
    use kimchi::{poly_commitment::commitment::CommitmentCurve, verifier::verify};
    use serde::Deserialize;

    use super::*;

    const KIMCHI_PROOF: &[u8] = include_bytes!("../kimchi_ec_add.proof");
    const KIMCHI_AGGREGATED_PUB_INPUT: &[u8] = include_bytes!("../kimchi_aggregated_pub_input.bin");

    #[test]
    fn kimchi_ec_add_proof_verifies() {
        let mut proof_buffer = [0u8; super::MAX_PROOF_SIZE];
        let proof_size = KIMCHI_PROOF.len();
        proof_buffer[..proof_size].clone_from_slice(KIMCHI_PROOF);

        let mut pub_input_buffer = [0u8; super::MAX_PUB_INPUT_SIZE];
        let pub_input_size = KIMCHI_AGGREGATED_PUB_INPUT.len();
        pub_input_buffer[..pub_input_size].clone_from_slice(KIMCHI_AGGREGATED_PUB_INPUT);

        let result =
            verify_kimchi_proof_ffi(&proof_buffer, proof_size, &pub_input_buffer, pub_input_size);

        assert!(result)
    }

    #[test]
    fn serialize_deserialize_pub_input_works() {
        let proof_file_path = Path::new("kimchi_ec_add.proof");
        let proof_file = std::fs::File::open(proof_file_path).expect("Could not open proof file");
        let proof_reader = BufReader::new(proof_file);
        let proof: ProverProof<Vesta, OpeningProof<Vesta>> =
            ProverProof::deserialize(&mut rmp_serde::Deserializer::new(proof_reader))
                .expect("Could not deserialize kimchi proof from file");

        let verifier_index_file_path = Path::new("kimchi_verifier_index.bin");
        let verifier_index_file = std::fs::File::open(verifier_index_file_path)
            .expect("Could not open verifier index file");
        let verifier_index_reader = BufReader::new(verifier_index_file);
        let mut verifier_index: VerifierIndex<Vesta, OpeningProof<Vesta>> =
            VerifierIndex::deserialize(&mut rmp_serde::Deserializer::new(verifier_index_reader))
                .expect("Could not deserialize verifier index");

        let srs_file_path = Path::new("kimchi_srs.bin");
        let srs_file = std::fs::File::open(srs_file_path).expect("Could not open SRS file");
        let srs_reader = BufReader::new(srs_file);
        let mut srs: SRS<Vesta> = SRS::deserialize(&mut rmp_serde::Deserializer::new(srs_reader))
            .expect("Could not deserialize verifier index");

        srs.add_lagrange_basis(verifier_index.domain);
        verifier_index.srs = Arc::new(srs.clone());
        verifier_index.endo = *Vesta::other_curve_endo();

        // sanity check that the proof verifies with the loaded files
        let group_map = <Vesta as CommitmentCurve>::Map::setup();
        assert!(
            verify::<Vesta, BaseSponge, ScalarSponge, OpeningProof<Vesta>>(
                &group_map,
                &verifier_index,
                &proof,
                &Vec::new(),
            )
            .is_ok()
        );

        // serialize and then deserialize aggregated kimchi pub inputs
        let pub_input_bytes = serialize_kimchi_pub_input(&verifier_index, &srs);
        let deserialized_verifier_index = deserialize_kimchi_pub_input(pub_input_bytes).unwrap();
        // verify the proof with the deserialized pub input (verifier index)
        assert!(
            verify::<Vesta, BaseSponge, ScalarSponge, OpeningProof<Vesta>>(
                &group_map,
                &deserialized_verifier_index,
                &proof,
                &Vec::new(),
            )
            .is_ok()
        );
    }
}
