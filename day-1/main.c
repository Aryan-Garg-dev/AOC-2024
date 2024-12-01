#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#define KB 1024
typedef unsigned int uint;

void parseString(char* string, int* num1, int* num2){
  char delimiter[] = "   ";
  char *token;
  token = strtok(string, delimiter);
  if (token != NULL) *num1 = atoi(token);
  token = strtok(NULL, delimiter);
  if (token != NULL) *num2 = atoi(token);

  printf("%d %d\n", *num1, *num2);
}

int max(int* arr, int len){
  int max = 0;
  for (int i = 0; i < len; i++){
    if (arr[i] > max) max = arr[i];
  }
  return max;
}

int* frequency(int* arr, int len){
  int* freqArr = (int*) malloc(sizeof(int) * (max(arr, len)+1));
  for (int i = 0; i < len; i++){
    freqArr[arr[i]] += 1;
  }
  return freqArr;
}

void merge(int* arr, uint left, uint mid, uint right){
  uint len_1 = mid - left + 1;
  uint len_2 = right - mid;

  int L[len_1], R[len_2];
  for (uint i = 0; i < len_1; i++) L[i] = arr[left + i];
  for (uint i = 0; i < len_1; i++) R[i] = arr[mid + i + 1];

  uint i = 0, j = 0, k = left;
  while (i < len_1 && j < len_2){
    if (L[i] <= R[j]) arr[k++] = L[i++];
    else arr[k++] = R[j++];
  }

  while (i < len_1) arr[k++] = L[i++];
  while (j < len_2) arr[k++] = R[j++];
}

void mergeSort(int* arr, uint left, uint right){
  if (left < right){
    uint mid = left + (right - left) / 2;
    mergeSort(arr, left, mid);
    mergeSort(arr, mid + 1, right);
    merge(arr, left, mid, right);
  }
}

int getDistance(int* arr1, int* arr2, uint len){
  int sum = 0;
  for (uint i = 0; i < len; i++){
    sum += (int) fabs((double) (arr1[i] - arr2[i]));
  }
  return sum;
}

int getSimilarityScore(int* arr1, int* arr2, int len){
  int* freqArr = frequency(arr2, len);
  int sum = 0;
  for (uint i = 0; i < len; i++) {
    sum += arr1[i] * freqArr[arr1[i]];
  }
  free(freqArr);
  return sum;
}

int main(){
  FILE* fptr;
  fptr = fopen("./input.txt", "r");

  char* input = (char*) malloc(sizeof(char)*KB);

  size_t arrlen = 1000;
  int num1[arrlen], num2[arrlen];
  uint i = 0, j = 0; 

  if (fptr != NULL){
    while (fgets(input, KB, fptr)){
      // puts(input);
      if(i < arrlen && j < arrlen) parseString(input, &num1[i++], &num2[j++]);
    }
  } 
  fclose(fptr);

  mergeSort(num1, 0, i-1);
  mergeSort(num2, 0, j-1);
  int distance = getDistance(num1, num2, arrlen);
  int similarityScore = getSimilarityScore(num1, num2, arrlen);
  printf("Distance: %d\n", distance);
  printf("Similarity Score: %d\n", similarityScore);

  return EXIT_SUCCESS;
}