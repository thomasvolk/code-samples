#!/usr/bin/env python3
from tkinter import *

class Mandelbrot(Frame):
    def __init__(self, master=None, **kw):
        super().__init__(master)
        self.pack()
        self.xmin = -2
        self.ymin = -2
        width = kw.get("width", 400)
        height = kw.get("height", 400)
        self.size = (width, height)
        self.canvas = Canvas(self, width = width, height = height)
        self.canvas.grid(row = 1, column = 1)
        self.rect = self.canvas.create_rectangle(0, 0, width, height, fill="#ffffff")
        self.canvas.pack()

    def draw(self, step, iterations):
        colorStep = 255.0 / iterations
        width, height = self.size
        for x in range(width):
            for y in range(height):
                point = complex(self.xmin + x * step, self.ymin + y * step)
                if abs(point) < 2:
                    nextpoint = complex(0, 0)
                    for i in range(iterations):
                        nextpoint = nextpoint * nextpoint + point
                        if abs(nextpoint) >= 2: break
                        color = self.getColor(iterations, colorStep, i)
                else:
                    color = "#000000"
                self._drawDot(x, y, color)

    def _drawDot(self, x, y, color):
        self.canvas.create_line(x, y, x+1, y+1, fill = color)
        

    def getColor(self, iterations, colorStep, i):
        # blue
        if i < iterations / 2:
            blue = 0
        else:
            blue = (i - iterations / 2) * colorStep
        # green
        if i > iterations / 2:
            green = 2 * (iterations -i) * colorStep
        else:
            green = 2 * i * colorStep
        # red
        if i > iterations / 2:
            red = 0
        else:
            red = 255 - 2 * i * colorStep;
        return "#%.2x%.2x%.2x" % (int(red), int(green), int(blue))


master = Tk()

m = Mandelbrot(master= master)
m.draw(0.01, 100)

mainloop()
