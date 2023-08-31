from PIL import Image
import numpy as np
from concurrent.futures
from timeit import default_timer as timer

# Determines the image pixel width and length
imageSize = int(input("Render Size : "))
imageSharpness = int(input(f"Render Sharpness [1 - {imageSize}]: "))
# Determines how precise the gradient is
maxIteration = imageSize / imageSharpness
# Default should be 2.5
valueRange = 2.5
# Default should be (-2, 1.25)
startingPoint = (-2, 1.25)

step = valueRange / (imageSize - 1)

inSetArray = np.zeros((imageSize, imageSize), dtype=np.float32)
colorArray = np.zeros((imageSize, imageSize, 3), dtype=np.float32)
colorArray[:] = 255

def mandelbrotTest(C): 
    Z = 0
    n = 0
    while abs(Z) <= 2 and n < maxIteration:
        Z = Z * Z + C
        n += 1
    return(n)

def inSetInterpret():
    for idx, x in np.ndenumerate(inSetArray):
        if inSetArray[idx] == 1:
            colorArray[idx] = (0, 0, 0)

def pointChecker(idx):
    complexNum = complex((startingPoint[0] + (idx[1] * step)), (startingPoint[1] - (idx[0] * step)))
    mandelbrotNum = mandelbrotTest(complexNum)
    if mandelbrotNum == maxIteration:
        inSetArray[idx] = 1

startTime = timer()

# https://www.youtube.com/watch?v=fKl2JW_qrso
# very helpful video

for idx, x in np.ndenumerate(inSetArray):
    print("stuff needs to go here but I don't have anything")
    # I need to pass idx into the lower function

inSetInterpret()

colorArray = colorArray.astype('uint8')
new_image = Image.fromarray(colorArray)
new_image.save('MandelbrotOptimized.png')

endTime = timer()
print("")
print(f"Values computed. Took [{endTime - startTime}] seconds.")