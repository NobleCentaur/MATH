from PIL import Image
import numpy as np
import random

imageWidth = 100
imageHeight = 100
maxItem = imageHeight * imageWidth
array = np.zeros((imageWidth, imageHeight, 3), dtype=np.float64)

def findLocation(Item, axis):
    if int(Item / imageWidth) <= imageHeight:
        #if the desired item location is below row 0
        if Item >= imageWidth :
            if axis == "x":
                return(int(Item / imageWidth))
            if axis == "y":
                return(Item%imageWidth)
        #if the desired item location is in row 0
        if Item < imageWidth:
            if axis == "x":
                return(0)
            if axis == "y":
                return(Item)

def noiseGeneration():
    for i in range(0, maxItem):
        array[findLocation(i, "x")][findLocation(i, "y")] = [random.randrange(0, 255), random.randrange(0, 255), random.randrange(0, 255)]

def verticalGradient(startColor, endColor):
    #Defines a list of the increment, pixel to pixel, of each rgb value
    stepX = np.zeros((1, 3), dtype=np.float64)
    stepX = (np.subtract(endColor, startColor)) / imageWidth
    #Set the first pixel of the first line that the loop will reference
    array[0][0] = startColor
    #sets the first line
    for x in range(0, imageHeight):
        array[x][0] = startColor
        for i in range(1, imageWidth):
            array[x][i] = np.add(array[x][i-1], stepX)

def diagonalGradient(startColor, endColor):
    #Defines a list of the increment, pixel to pixel, of each rgb value
    stepX = np.zeros((1, 3), dtype=np.float64)
    stepY = np.zeros((1, 3), dtype=np.float64)
    stepX = np.subtract(endColor, startColor) / imageWidth
    stepY = np.subtract(endColor, startColor) / imageHeight
    #Set the first pixel of the first line that the loop will reference
    array[0][0] = startColor
    #sets the first line
    for x in range(0, imageHeight):
        array[x][0] = np.add(array[x-1][0], stepY)
        for i in range(1, imageWidth):
            array[x][i] = [(array[x][i-1][0] + stepX[0]), (array[x][i-1][1] + stepX[1]), (array[x][i-1][2] + stepX[2])]

verticalGradient((15, 254, 0), (12, 24, 124))
#diagonalGradient((0, 255, 255), (255, 0, 255))

array = array.astype('uint8')

new_image = Image.fromarray(array)
new_image.save('Image.png')