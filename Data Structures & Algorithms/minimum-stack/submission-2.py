class MinStack:

    def __init__(self):
        self.stack = []
        self.length = 0
        self.Min = []

    def push(self, val: int) -> None:
        self.stack.append(val)
        self.length += 1
        self.Min.append(val)
        self.Min.sort()

    def pop(self) -> None:
        if self.length >= 1:
            elem = self.stack.pop()
            self.length -=1
            self.Min.remove(elem)


    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.Min[0]
