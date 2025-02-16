use rayt::vec3::Vec3;

#[cfg(test)]
use assert_approx_eq::assert_approx_eq;

#[test]
fn test_gen() {
    let p = Vec3 {
        x: 0.1,
        y: 0.2,
        z: 0.3,
    };
    assert_eq!(p.x(), 0.1);
    assert_eq!(p.y(), 0.2);
    assert_eq!(p.z(), 0.3);

    let q = Vec3::new(0.2, 0.3, 0.4);
    assert_eq!(q.x(), 0.2);
    assert_eq!(q.y(), 0.3);
    assert_eq!(q.z(), 0.4);
}

#[test]
fn test_add() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    let q = Vec3::new(0.3, 0.4, 0.5);
    let r = p + q;

    assert_approx_eq!(r.x(), 0.3);
    assert_approx_eq!(r.y(), 0.9);
    assert_approx_eq!(r.z(), 1.2);
}

#[test]
fn test_sub() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    let q = Vec3::new(0.2, 0.3, 0.4);
    let r = p - q;
    assert_approx_eq!(r.x(), -0.1);
    assert_approx_eq!(r.y(), -0.1);
    assert_approx_eq!(r.z(), -0.1);
}

#[test]
fn test_neg() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    let q = -p;
    assert_approx_eq!(q.x(), -0.1);
    assert_approx_eq!(q.y(), -0.2);
    assert_approx_eq!(q.z(), -0.3);
}

#[test]
fn test_mul() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    let q = Vec3::new(0.2, 0.3, 0.4);
    let r = p * q;
    assert_approx_eq!(r.x(), 0.02);
    assert_approx_eq!(r.y(), 0.06);
    assert_approx_eq!(r.z(), 0.12);
}

#[test]
fn test_div() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    let q = Vec3::new(0.2, 0.3, 0.4);
    let r = p / q;
    assert_approx_eq!(r.x(), 0.5);
    assert_approx_eq!(r.y(), 0.6666666666666666);
    assert_approx_eq!(r.z(), 0.3 / 0.4);
}

#[test]
fn test_dot() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    let q = Vec3::new(0.2, 0.3, 0.4);
    assert_approx_eq!(p.dot(&q), 0.2);
}

#[test]
fn test_length_squared() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    assert_approx_eq!(p.length_squared(), 0.14);
}

#[test]
fn test_random() {
    let p = Vec3::random(-1.0, 1.0);
    assert!(p.x() >= -1.0 && p.x() <= 1.0);
    assert!(p.y() >= -1.0 && p.y() <= 1.0);
    assert!(p.z() >= -1.0 && p.z() <= 1.0);
}

#[test]
fn test_near_zero() {
    let p = Vec3::new(0.1, 0.2, 0.3);
    assert!(!p.near_zero());
    let p = Vec3::new(0.0, 0.0, 0.0);
    assert!(p.near_zero());
}
