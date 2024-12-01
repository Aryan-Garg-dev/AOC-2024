#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#define KB 1024
#define STRING_SIZE 128
typedef unsigned int uint;

void parseString(char* string, int* num1, int* num2);
int max(int* arr, size_t len);
int* frequency(int* arr, size_t len);
void merge(int* arr, uint left, uint mid, uint right);
void mergeSort(int* arr, uint left, uint right);
int getDistance(int* arr1, int* arr2, size_t len);
int getSimilarityScore(int* arr1, int* arr2, size_t len);

int main(){
  FILE* fptr;
  fptr = fopen("./input.txt", "r");

  char* input = (char*) malloc(sizeof(char)*STRING_SIZE);
  size_t arrlen = 1000;
  int num1[arrlen], num2[arrlen];
  uint i = 0; 

  if (fptr != NULL){
    while (fgets(input, STRING_SIZE, fptr)){
      if(i < arrlen) {
        parseString(input, &num1[i], &num2[i]);
        i++;
      }
    }
  } 
  fclose(fptr);
  free(input);

  mergeSort(num1, 0, i-1);
  mergeSort(num2, 0, i-1);

  int distance = getDistance(num1, num2, i);
  int similarityScore = getSimilarityScore(num1, num2, i);
  printf("Distance: %d\n", distance);
  printf("Similarity Score: %d\n", similarityScore);

  return EXIT_SUCCESS;
}

void parseString(char* string, int* num1, int* num2){
  char delimiter[] = "   ";
  char *token;
  token = strtok(string, delimiter);
  if (token != NULL) *num1 = atoi(token);
  token = strtok(NULL, delimiter);
  if (token != NULL) *num2 = atoi(token);
}

int max(int* arr, size_t len){
  if (len == 0) return -1;
  int max = 0;
  for (uint i = 0; i < len; i++){
    if (arr[i] > max) max = arr[i];
  }
  return max;
}

int* frequency(int* arr, size_t len){
  int* freqArr = (int*) calloc((max(arr, len)+1), sizeof(int));
  for (uint i = 0; i < len; i++){
    freqArr[arr[i]] += 1;
  }
  return freqArr;
}

void merge(int* arr, uint left, uint mid, uint right){
  uint len_1 = mid - left + 1;
  uint len_2 = right - mid;

  int L[len_1], R[len_2];
  for (uint i = 0; i < len_1; i++) L[i] = arr[left + i];
  for (uint i = 0; i < len_2; i++) R[i] = arr[mid + i + 1];

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

int getDistance(int* arr1, int* arr2, size_t len){
  int sum = 0;
  for (uint i = 0; i < len; i++){
    sum += abs((arr1[i] - arr2[i]));
  }
  return sum;
}

int getSimilarityScore(int* arr1, int* arr2, size_t len){
  int* freqArr = frequency(arr2, len);
  int sum = 0;
  for (uint i = 0; i < len; i++) {
    sum += arr1[i] * freqArr[arr1[i]];
  }
  free(freqArr);
  return sum;
}
