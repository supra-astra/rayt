use std::ops::{
    Add, AddAssign, Div, DivAssign, Index, IndexMut, Mul, MulAssign, Neg, Sub, SubAssign,
};

use rand::Rng;

//taken from here:
// https://github.com/ryankaplan/vec3/blob/master/src/lib.rs
#[derive(Copy, Clone, Debug, Deserialize, Serialize)]
pub struct Vec3 {
    pub x: f64,
    pub y: f64,
    pub z: f64,
}

impl Vec3 {
    pub fn new(e0: f64, e1: f64, e2: f64) -> Vec3 {
        Vec3 {
            x: e0,
            y: e1,
            z: e2,
        }
    }
    pub fn x(&self) -> f64 {
        self.x
    }

    pub fn y(&self) -> f64 {
        self.y
    }

    pub fn z(&self) -> f64 {
        self.z
    }

    pub fn r(&self) -> f64 {
        self.x
    }

    pub fn g(&self) -> f64 {
        self.y
    }
    pub fn b(&self) -> f64 {
        self.z
    }

    pub fn zero() -> Vec3 {
        Vec3::new(0.0, 0.0, 0.0)
    }

    pub fn one() -> Vec3 {
        Vec3::new(1.0, 1.0, 1.0)
    }

    //vector operations
    pub fn length(&self) -> f64 {
        self.length_squared().sqrt()
    }

    pub fn length_squared(&self) -> f64 {
        self.x * self.x + self.y * self.y + self.z * self.z
    }

    pub fn dot(&self, other: &Vec3) -> f64 {
        self.x * other.x + self.y * other.y + self.z * other.z
    }

    pub fn cross(&self, other: &Vec3) -> Vec3 {
        Vec3::new(
            self.y * other.z - self.z * other.y,
            self.z * other.z - self.x * other.z,
            self.x * other.y - self.y * other.x,
        )
    }

    pub fn unit_vector(&self) -> Vec3 {
        *self / self.length()
    }

    //useful utilities

    pub fn min(&self, other: &Vec3) -> Vec3 {
        Vec3::new(
            self.x.min(other.x),
            self.y.min(other.y),
            self.z.min(other.z),
        )
    }

    pub fn max(&self, other: &Vec3) -> Vec3 {
        Vec3::new(
            self.x.max(other.x),
            self.y.max(other.y),
            self.z.max(other.z),
        )
    }

    pub fn clamp(&self, min: &Vec3, max: &Vec3) -> Vec3 {
        self.max(min).min(max)
    }

    pub fn lerp(&self, other: &Vec3, t: f64) -> Vec3 {
        (1.0 - t) * *self + t * *other
    }

    pub fn apply<T>(&self, f: T) -> Vec3
    where
        T: Fn(f64) -> f64,
    {
        Vec3::new(f(self.x), f(self.y), f(self.z))
    }

    pub fn random(min: f64, max: f64) -> Vec3 {
        let mut rng = rand::thread_rng();
        Vec3::new(
            rng.gen_range(min..max),
            rng.gen_range(min..max),
            rng.gen_range(min..max),
        )
    }

    pub fn random_in_unit_sphere() -> Vec3 {
        loop {
            let p = Vec3::random(-1.0, 1.0);
            if p.length_squared() < 1.0 {
                return p;
            }
        }
    }

    pub fn distance(&self, other: &Vec3) -> f64 {
        let dx = self.x - other.x();
        let dy = self.y - other.y();
        let dz = self.z - other.z();

        (dx * dx + dy * dy + dz * dz).sqrt()
    }
}

//defining the operators for different operators for our vector class

impl IndexMut<usize> for Vec3 {
    fn index_mut(&mut self, idx: usize) -> &mut Self::Output {
        match idx {
            0 => &mut self.x,
            1 => &mut self.y,
            2 => &mut self.z,
            _ => panic!("Index out of bounds for Vec3"),
        }
    }
}

//index at
impl Index<usize> for Vec3 {
    type Output = f64;

    fn index(&self, idx: usize) -> &Self::Output {
        match idx {
            0 => &self.x,
            1 => &self.y,
            2 => &self.z,
            _ => panic!("Index out of bounds for Vec3: {}", idx),
        }
    }
}

//negative
impl Neg for Vec3 {
    type Output = Vec3;

    fn neg(self) -> Self::Output {
        Vec3::new(-self.x, -self.y, -self.z)
    }
}

//add
impl Add for Vec3 {
    type Output = Vec3;

    fn add(self, other: Self) -> Self::Output {
        Vec3::new(self.x + other.x, self.y + other.y, self.z + other.z)
    }
}

//+=
impl AddAssign for Vec3 {
    fn add_assign(&mut self, other: Self) {
        self.x += other.x;
        self.y += other.y;
        self.z += other.z;
    }
}

//-
impl Sub for Vec3 {
    type Output = Vec3;

    fn sub(self, other: Self) -> Self::Output {
        Vec3::new(self.x - other.x, self.y - other.y, self.z - other.z)
    }
}

//-=
impl SubAssign for Vec3 {
    fn sub_assign(&mut self, other: Self) {
        self.x -= other.x;
        self.y -= other.y;
        self.z -= other.z;
    }
}

// mul for floats
impl Mul<f64> for Vec3 {
    type Output = Vec3;

    fn mul(self, scalar: f64) -> Self::Output {
        Vec3::new(self.x * scalar, self.y * scalar, self.z * scalar)
    }
}

//mul for vec3
impl Mul<Vec3> for f64 {
    type Output = Vec3;
    fn mul(self, vec: Vec3) -> Self::Output {
        vec * self
    }
}

// *=

impl MulAssign<f64> for Vec3 {
    fn mul_assign(&mut self, scalar: f64) {
        self.x *= scalar;
        self.y *= scalar;
        self.z *= scalar;
    }
}

// /
impl Div<f64> for Vec3 {
    type Output = Vec3;

    fn div(self, scalar: f64) -> Self::Output {
        Vec3::new(self.x / scalar, self.y / scalar, self.z / scalar)
    }
}

// /=
impl DivAssign<f64> for Vec3 {
    fn div_assign(&mut self, scalar: f64) {
        let k = 1.0 / scalar;
        self.x *= k;
        self.y *= k;
        self.z *= k;
    }
}

//element wise multiplication
impl Mul<Vec3> for Vec3 {
    type Output = Vec3;

    fn mul(self, other: Vec3) -> Self::Output {
        Vec3::new(self.x * other.x, self.y * other.y, self.z * other.z)
    }
}

//element wise division
impl Div<Vec3> for Vec3 {
    type Output = Vec3;

    fn div(self, other: Vec3) -> Self::Output {
        Vec3::new(self.x / other.x, self.y / other.y, self.z / other.z)
    }
}

//other operations

pub fn dot(a: &Vec3, b: &Vec3) -> f64 {
    a.dot(b)
}

pub fn cross(a: &Vec3, b: &Vec3) -> Vec3 {
    a.cross(b)
}

pub fn unit_vector(v: &Vec3) -> Vec3 {
    v.unit_vector()
}
