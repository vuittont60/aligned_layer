use sp1_verifier_wrapper::verify_sp1_proof_ffi;
use sp1_verifier_wrapper::MAX_PROOF_SIZE;

const PROOF: &[u8; 1040380] = include_bytes!("../../../../tests/testing_data/sp1-fibonacci.proof");

fn main() {
    let mut proof_buffer = [0u8; MAX_PROOF_SIZE];
    let proof_size = PROOF.len();
    proof_buffer[..proof_size].clone_from_slice(PROOF);
    let result = verify_sp1_proof_ffi(&proof_buffer, proof_size);
    assert!(result)
}
