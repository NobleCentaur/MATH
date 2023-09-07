from PIL import Image
import numpy as np
import random

imageWidth = 100
imageHeight = 100
maxItem = imageHeight * imageWidth
row = 0
column = 0
currentItemLocation = [row, column]
image = []


for i in range(imageHeight):
    image.append([])
for i in range(len(image)):
    image[i] = [[0, 0, 0] for j in range(imageWidth)]

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

# Generates completely random image
def noiseGeneration():
    for i in range(0, maxItem):
        image[findLocation(i, "x")][findLocation(i, "y")] = [random.randrange(0, 255), random.randrange(0, 255), random.randrange(0, 255)]

#def gradient(startColor, endColor):
#    for i in range(0, imageWidth):
#        if image[0][i][0] < 255:
#            image[0][i] = [image[0][i][0] + (5*i), 0, 0]
#    for i in range(0, imageHeight):
#        image[i][0] = [image[i][0], 0, 0]

#heeeee eheeeeeee comment goes here
noiseGeneration()

array = np.array(image, dtype=np.uint8)

new_image = Image.fromarray(array)
new_image.save('Image.png')