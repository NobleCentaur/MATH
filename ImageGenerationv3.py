from PIL import Image
import numpy as np
import random
import cmath

imageWidth = 100
imageHeight = 100

if imageHeight % 2 == 0:
    imageHeight =  imageHeight + 1
if imageWidth % 2 == 0:
    imageWidth =  imageWidth + 1

array = np.zeros((imageWidth, imageHeight, 3), dtype=np.float64)
for x in range(0, imageHeight):
    for i in range(0, imageWidth):
        array[x][i] = [255, 255, 255]

def mandelbrotTest(C):
    Z = [0]
    for i in range(6):
        Z.append(Z[-1] * Z[-1] + C)
    print(Z)

mandelbrotTest(-1+0.25j)

array = array.astype('uint8')

new_image = Image.fromarray(array)
new_image.save('Image.png')