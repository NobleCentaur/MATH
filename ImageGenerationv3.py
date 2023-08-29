from PIL import Image
import numpy as np
import random

imageSize = 4
maxIteration = 80
step = 3 / (imageSize - 1)

# (imageSize)x = 3

array = np.zeros((imageSize, imageSize, 3), dtype=np.float64)
for x in range(0, imageSize):
    for i in range(0, imageSize):
        array[x][i] = [255, 255, 255]

def mandelbrotTest(C): 
    Z = 0
    n = 0
    while abs(Z) <= 2 and n < maxIteration:
        Z = Z * Z + C
        n += 1
    return(n)

# input 0,0 central coordinate plane coordinate and returns a 0,0 top left plane coordinate
def coordsToArray(x, y):
    arrayOffset = (imageSize - 1) / 2  
    returnX = np.abs(y - arrayOffset)
    returnY = x + arrayOffset
    return(returnX, returnY)



#for a in range(imageSize):
#    for b in range(imageSize):
#        complexNum = complex((-2 + (a * step)), (1.5 - (b * step)))
#        print(mandelbrotTest(complexNum))
#        if mandelbrotTest(complexNum) == True:
#            array[b][a] = (0, 0, 0)

#var1 = 0
#for i in range(0, imageSize):
#    var1 = -2 + (i * step)  
#    print(var1)
#for i in range(imageSize):
#    var1 = 1.5 - (i * step)
#    print(var1)

array = array.astype('uint8')

new_image = Image.fromarray(array)
new_image.save('Image.png')