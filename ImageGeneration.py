from PIL import Image
import numpy as np
import random

imageWidth = 10
imageHeight = 10
currentItem = 0
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
for i in range(0, maxItem):
    image[findLocation(i, "x")][findLocation(i, "y")] = [random.randrange(0, 255), random.randrange(0, 255), random.randrange(0, 255)]

var1 = 10
for i in range(0, imageWidth):
    image[0][i] = [image[0][i] + (5*i), 0, 0]

var2 = image[0][1]
for i in range(0, imageHeight):
    image[i][0] = [image[i][0], 0, 0]

array = np.array(image, dtype=np.uint8)

new_image = Image.fromarray(array)
new_image.save('Image.png')