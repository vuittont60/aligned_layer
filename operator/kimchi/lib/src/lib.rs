// #[no_mangle]
// pub extern "C" fn verify_kimchi_proof_ffi(
//     proof_bytes: &[u8; MAX_PROOF_SIZE],
//     proof_len: usize,
// ) -> bool {
//     if let Ok(proof) = bincode::deserialize(&proof_bytes[..proof_len]) {
//         // ...
//         // add kimchi verifier code here
//         // ...
//     }

//     false
// }

// #[cfg(test)]
// mod tests {
//     use super::*;

//     #[test]
//     fn it_works() {
//         let result = add(2, 2);
//         assert_eq!(result, 4);
//     }
// }
