from PIL import Image
import numpy as np
import multiprocessing
from timeit import default_timer as timer

# Determines the image pixel width and length
imageSize = 10
imageSharpness = 1
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

#def mandelbrotTest(C): 
#    Z = 0
#    n = 0
#    while abs(Z) <= 2 and n < maxIteration:
#        Z = Z * Z + C
#        n += 1
#    return(n)

def pointChecker(idx):
    complexNum = complex((startingPoint[0] + (idx[1] * step)), (startingPoint[1] - (idx[0] * step)))
    Z = 0
    n = 0
    while abs(Z) <= 2 and n < maxIteration:
        Z = Z * Z + complexNum
        n += 1
    print(n)
    if n == maxIteration:
        print("YAY")
        return(idx)
    else:
        return(0)

if __name__ == "__main__":

    startTime = timer()
    coordinates = []
    for x, _ in np.ndenumerate(inSetArray):
        coordinates.append(x)

    # https://www.youtube.com/watch?v=fKl2JW_qrso
    # very helpful video

    processes = []
    for i in range(imageSize**2):
        p = multiprocessing.Process(target=pointChecker, args=[coordinates[i]])
        p.start()
        processes.append(p)

    for process in processes:
        process.join()

    #interprets 2d binary array to color array
    for idx, x in np.ndenumerate(inSetArray):
        if inSetArray[idx] == 1:
            colorArray[idx] = (0, 0, 0)

    colorArray = colorArray.astype('uint8')
    new_image = Image.fromarray(colorArray)
    new_image.save('MandelbrotOptimized.png')

    endTime = timer()
    print("")
    print(f"Values computed. Took [{endTime - startTime}] seconds.")