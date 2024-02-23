use sp1_core::{utils::BabyBearBlake3, SP1ProofWithIO, SP1Verifier};

const ELF: &[u8] = include_bytes!("../elf/riscv32im-succinct-zkvm-elf");
pub const MAX_PROOF_SIZE: usize = 1024 * 1024;

#[no_mangle]
pub extern "C" fn verify_sp1_proof_ffi(
    proof_bytes: &[u8; MAX_PROOF_SIZE],
    proof_len: usize,
) -> bool {
    let proof: SP1ProofWithIO<BabyBearBlake3> =
        bincode::deserialize(&proof_bytes[..proof_len]).expect("Could not open proof file");

    SP1Verifier::verify(ELF, &proof).is_ok()
}

#[cfg(test)]
mod tests {
    use super::*;

    const PROOF: &[u8; 1040380] =
        include_bytes!("../../../../tests/testing_data/sp1-fibonacci.proof");

    #[test]
    fn verify_sp1_proof_works() {
        let mut proof_buffer = [0u8; super::MAX_PROOF_SIZE];
        let proof_size = PROOF.len();
        proof_buffer[..proof_size].clone_from_slice(PROOF);
        let result = verify_sp1_proof_ffi(&proof_buffer, proof_size);
        assert!(result)
    }
}
