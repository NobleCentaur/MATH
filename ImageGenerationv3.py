from PIL import Image
import numpy as np
import random

imageSize = 5
maxIteration = 20
imageShift = 0

if imageSize % 2 == 0:
    imageSize =  imageSize + 1

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

# arrayToCoords before
# [(0, 0) , (0, 1) , (0, 2) , (0, 3) , (0, 4)]
# [(1, 0) , (1, 1) , (1, 2) , (1, 3) , (1, 4)]
# [(2, 0) , (2, 1) , (2, 2) , (2, 3) , (2, 4)]
# [(3, 0) , (3, 1) , (3, 2) , (3, 3) , (3, 4)]
# [(4, 0) , (4, 1) , (4, 2) , (4, 3) , (4, 4)]
# convert from locations on the coordinate plane to this
# [(-2, 2) , (-1, 2) , (0, 1) , (1, 2) , (2, 2) ] 
# [(-2, 1) , (-1, 1) , (0, 1) , (1, 1) , (2, 1) ]
# [(-2, 0) , (-1, 0) , (0, 0) , (1, 0) , (2, 0) ]
# [(-2, -1), (-1, -1), (0, -1), (1, -1), (2, -1)]
# [(-2, -2), (-1, -2), (0, -2), (1, -2), (2, -2)]

#def coordsToArray(x, y):
#    arrayOffset = (imageSize - 1) / 2  
#    return(x, y)

#print(coordsToArray(2, -2))

array = array.astype('uint8')

new_image = Image.fromarray(array)
new_image.save('Image.png')