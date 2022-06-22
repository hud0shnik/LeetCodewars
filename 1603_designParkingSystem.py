#!/usr/bin/python3

class ParkingSystem:

    def __init__(self, big: int, medium: int, small: int):
        self.nums = {1:big, 2:medium, 3:small}

    def addCar(self, carType: int) -> bool:
        self.nums[carType] -= 1
        return self.nums[carType] >= 0

if __name__ == "__main__":
    p = ParkingSystem(2, 1, 0)
    print(p.addCar(1))
    print(p.addCar(1))
    print(p.addCar(1))
    print(p.addCar(2))
    print(p.addCar(3))
