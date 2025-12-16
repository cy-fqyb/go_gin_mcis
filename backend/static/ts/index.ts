type Compare<T> = (a: T, b: T) => number;

const defaultCompare = <T>(a: T & (number | string), b: T & (number | string)): number => {
    // 通用数字/字符串比较器
    if (a < b) return -1;
    if (a > b) return 1;
    return 0;
};

class Sorter {
    // 计时器：返回结果与耗时，并支持标签
    static executeAndTime<T>(fn: () => T, label: string = ''): { result: T; seconds: number } {
        const start = performance.now();
        const result = fn();
        const end = performance.now();
        const timeInSeconds = (end - start) / 1000;
        const msg = label ? `[${label}] Execution time: ${timeInSeconds.toFixed(6)} s` : `Execution time: ${timeInSeconds.toFixed(6)} s`;
        console.log(msg);
        return { result, seconds: timeInSeconds };
    }

    // 快速排序（不可变返回）：使用 median-of-three + 小段插入排序，避免最坏退化
    static quickSort<T>(arr: T[], compare: Compare<T> = defaultCompare as Compare<T>): T[] {
        const a = arr.slice(); // 保持不可变
        if (a.length <= 1) return a;
        // 小数组时改用插入排序
        const threshold = 16;

        const swap = (i: number, j: number) => {
            const tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        };

        const insertionSortRange = (lo: number, hi: number) => {
            for (let i = lo + 1; i <= hi; i++) {
                const key = a[i];
                let j = i - 1;
                while (j >= lo && compare(a[j], key) > 0) {
                    a[j + 1] = a[j];
                    j--;
                }
                a[j + 1] = key;
            }
        };

        const medianOfThreeIndex = (i: number, j: number, k: number): number => {
            const A = a[i], B = a[j], C = a[k];
            // 返回中位数所在的索引
            if (compare(A, B) < 0) {
                if (compare(B, C) < 0) return j; // A < B < C
                return compare(A, C) < 0 ? k : i; // A < C <= B or C <= A < B
            } else {
                if (compare(A, C) < 0) return i; // B <= A < C
                return compare(B, C) < 0 ? k : j; // B < C <= A or C <= B <= A
            }
        };

        const quickSortInPlace = (lo: number, hi: number) => {
            while (lo < hi) {
                if (hi - lo + 1 <= threshold) {
                    insertionSortRange(lo, hi);
                    return;
                }

                const mid = lo + ((hi - lo) >> 1);
                const pivotIndex = medianOfThreeIndex(lo, mid, hi);
                // Lomuto 分区：将 pivot 放到末尾
                swap(pivotIndex, hi);
                const pivot = a[hi];

                let i = lo;
                for (let j = lo; j < hi; j++) {
                    if (compare(a[j], pivot) < 0) {
                        swap(i, j);
                        i++;
                    }
                }
                swap(i, hi);

                // 先处理较小的区间以限制递归深度（转换为迭代的尾递归优化）
                const leftSize = i - 1 - lo + 1; // = i - lo
                const rightSize = hi - (i + 1) + 1; // = hi - i
                if (leftSize < rightSize) {
                    if (lo < i - 1) quickSortInPlace(lo, i - 1);
                    lo = i + 1; // 尾递归优化处理右侧
                } else {
                    if (i + 1 < hi) quickSortInPlace(i + 1, hi);
                    hi = i - 1; // 尾递归优化处理左侧
                }
            }
        };

        quickSortInPlace(0, a.length - 1);
        return a;
    }

    // 冒泡排序（不可变返回，稳定）
    static bubbleSort<T>(arr: T[], compare: Compare<T> = defaultCompare as Compare<T>): T[] {
        const a = arr.slice();
        const len = a.length;
        for (let i = 0; i < len - 1; i++) {
            let swapped = false;
            for (let j = 0; j < len - 1 - i; j++) {
                if (compare(a[j], a[j + 1]) > 0) {
                    [a[j], a[j + 1]] = [a[j + 1], a[j]];
                    swapped = true;
                }
            }
            if (!swapped) break; // 已有序，提前结束
        }
        return a;
    }

    // 插入排序（不可变返回，稳定）
    static insertionSort<T>(arr: T[], compare: Compare<T> = defaultCompare as Compare<T>): T[] {
        const a = arr.slice();
        const len = a.length;
        for (let i = 1; i < len; i++) {
            const key = a[i];
            let j = i - 1;
            while (j >= 0 && compare(a[j], key) > 0) {
                a[j + 1] = a[j];
                j--;
            }
            a[j + 1] = key;
        }
        return a;
    }

    // 选择排序（不可变返回，非稳定）
    static selectionSort<T>(arr: T[], compare: Compare<T> = defaultCompare as Compare<T>): T[] {
        const a = arr.slice();
        const len = a.length;
        for (let i = 0; i < len - 1; i++) {
            let minIndex = i;
            for (let j = i + 1; j < len; j++) {
                if (compare(a[j], a[minIndex]) < 0) minIndex = j;
            }
            if (minIndex !== i) [a[i], a[minIndex]] = [a[minIndex], a[i]];
        }
        return a;
    }
}

// 示例用法
const arrayToSort = [64, 34, 25, 12, 22, 11, 90];

Sorter.executeAndTime(() => Sorter.quickSort(arrayToSort), 'QuickSort');
Sorter.executeAndTime(() => Sorter.bubbleSort(arrayToSort), 'BubbleSort');
// 也可测试插入与选择
Sorter.executeAndTime(() => Sorter.insertionSort(arrayToSort), 'InsertionSort');
Sorter.executeAndTime(() => Sorter.selectionSort(arrayToSort), 'SelectionSort');

