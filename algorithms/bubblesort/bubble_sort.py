def bubble_sort(array):
    for i in array:
        for j in range(len(array) - 1):
            if array[j] > array[j + 1]:
                array[j], array[j+1] = array[j+1], array[j]
    return array

def main():
    unsorted = [9, 2, 8, 3, 1, 10, 6, 7, 4, 5]
    print(bubble_sort(unsorted))

if __name__ == "__main__":
    main()
