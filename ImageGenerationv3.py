from PIL import Image
import numpy as np
import random
import cmath

imageWidth = 100
imageHeight = 100
maxIteration = 20
imageShift = 0

if imageHeight % 2 == 0:
    imageHeight =  imageHeight + 1
if imageWidth % 2 == 0:
    imageWidth =  imageWidth + 1

array = np.zeros((imageWidth, imageHeight, 3), dtype=np.float64)
for x in range(0, imageHeight):
    for i in range(0, imageWidth):
        array[x][i] = [255, 255, 255]

def mandelbrotTest(C):
    Z = 0
    n = 0
    while abs(Z) <= 2 and n < maxIteration:
        Z = Z * Z + C
        n += 1
    return(n)

def arrayToCoords(x, y):
    centerX = ((imageWidth - 1) / 2) + 1
    centerY = ((imageHeight - 1) / 2) + 1
    
    return()

print(mandelbrotTest(-1+0.25j))

array = array.astype('uint8')

new_image = Image.fromarray(array)
new_image.save('Image.png')