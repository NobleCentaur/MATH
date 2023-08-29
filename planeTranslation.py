import numpy as np

# IMPORTANT
# only the array thing works
# I figured it out in another document

imageSize = 5
array = np.zeros((imageSize, imageSize, 2), dtype=np.float64)

# arrayToCoords before
# [(1, 1) , (1, 2) , (1, 3) , (1, 4) , (1, 5)]
# [(2, 1) , (2, 2) , (2, 3) , (2, 4) , (2, 5)]
# [(3, 1) , (3, 2) , (3, 3) , (3, 4) , (3, 5)]
# [(4, 1) , (4, 2) , (4, 3) , (4, 4) , (4, 5)]
# [(5, 1) , (5, 2) , (5, 3) , (5, 4) , (5, 5)]
# convert from locations on the coordinate plane to this
# [(-2, 2) , (-1, 2) , (0, 1) , (1, 2) , (2, 2) ] 
# [(-2, 1) , (-1, 1) , (0, 1) , (1, 1) , (2, 1) ]
# [(-2, 0) , (-1, 0) , (0, 0) , (1, 0) , (2, 0) ]
# [(-2, -1), (-1, -1), (0, -1), (1, -1), (2, -1)]
# [(-2, -2), (-1, -2), (0, -2), (1, -2), (2, -2)]

def arrayPlaneTranslation():
    smallestX = 0
    largestY = 0
    for i in range(imageSize):
        for j in range(imageSize):
            if array[i][j][0] < smallestX:
                smallestX = array[i][j][0]
            if array[i][j][1] > largestY:
                largestY = array[i][j][1]
    smallestX = np.abs(smallestX)
    print(smallestX)
    print(largestY)
    for i in range(imageSize):
        for j in range(imageSize):
            array[i][j][0] = array[i][j][0] + smallestX
            array[i][j][1] = abs(array[i][j][1] - largestY)

def planeTranslation(x, y):
    y = x + imageSize
    x = abs(y - imageSize)
    return(x, y)

print(planeTranslation(2, 2))