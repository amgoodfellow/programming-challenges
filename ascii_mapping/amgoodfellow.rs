fn calc_num(word : &str) -> f32 {
    let total : usize = word.chars().map(|x| {x as usize}).sum();
    total as f32 / word.len() as f32
}

fn main() {
    println!("{:?}", calc_num("Hello"));
    println!("{:?}", calc_num("Test"));
    println!("{:?}", calc_num("QWERTY"));
    println!("{:?}", calc_num("!@#$%^&"));
    println!("{:?}", calc_num("This is fun!"));
}
