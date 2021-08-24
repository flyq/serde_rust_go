use serde::Serialize;

#[derive(Serialize)]
struct SignatureContent {
    token_type: String,
    price: u64,
    time_stamp: u64,
}

fn main() {
    let sig_content = SignatureContent {
        token_type: "ETH".to_string(),
        price: 10000,
        time_stamp: 0,
    };
    let msgbuf = serde_json::to_vec(&sig_content).unwrap();
    println!("msg bytes: {:?}\nmsg in hex: {:?}\n", msgbuf, hex::encode(&msgbuf));
}
