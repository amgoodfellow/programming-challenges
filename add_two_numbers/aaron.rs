fn helper(a: u8, b: u8, carry: u8) -> (u8, u8) {
    let total = a + b + carry;

    if total >= 10 {
        return (total % 10, 1);
    }

    (total, 0)
}

fn add(one: Vec<u8>, two: Vec<u8>) -> Vec<u8> {
    let mut iter_one = one.iter();
    let mut iter_two = two.iter();

    let mut result: Vec<u8> = Vec::new();
    let mut carry = 0;

    loop {
        match (iter_one.next(), iter_two.next()) {
            (Some(a), Some(b)) => {
                let (res, car) = helper(*a, *b, carry);
                result.push(res);
                carry = car;
            }
            (Some(a), None) => {
                let (res, car) = helper(*a, 0, carry);
                result.push(res);
                carry = car;
            }
            (None, Some(b)) => {
                let (res, car) = helper(0, *b, carry);
                result.push(res);
                carry = car;
            }
            _ => break,
        }
    }
    result
}

//To compile `rustc aaron.rs`
fn main() {
    assert_eq!(add(vec![2, 4, 3], vec![5, 6, 4]), vec![7, 0, 8]);
    assert_eq!(add(vec![2, 4, 3], vec![5]), vec![7, 4, 3]);
    assert_eq!(
        add(
            vec![3, 2, 2, 2, 5, 0, 4, 8, 6],
            vec![6, 0, 4, 8, 4, 6, 7, 7, 5, 1, 4, 5, 6, 7, 8]
        ),
        vec![9, 2, 6, 0, 0, 7, 1, 6, 2, 2, 4, 5, 6, 7, 8]
    );
}
