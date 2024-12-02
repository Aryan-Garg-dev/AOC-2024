import java.io.*;
import java.util.*;

public class Main {
  public static List<Integer> parseInput(String input){
    List<Integer> list = new ArrayList<>();
    String[] numStrings =  input.split(" ");
    for (String num: numStrings){
      list.add(Integer.valueOf(num));
    }
    return list;
  }

  public static boolean validateSafe(List<Integer> input){
    int prev = input.get(0);
    int prevDiff = input.get(1) - input.get(0);
    for (int i = 1; i < input.size(); i++){
      int curr = input.get(i);
      int diff = curr - prev;
      int absDiff = Math.abs(diff);
      if (absDiff != 1 && absDiff != 2 && absDiff != 3) return false;
      if (diff * prevDiff < 0) return false;
      prev = curr;
      prevDiff = diff;
    }
    return true;
  }

  public static boolean isSafeAfterRemBadLevel(List<Integer> input){
    if (validateSafe(input)) return true;
    for (int i = 0; i < input.size(); i++){
      List<Integer> temp = new ArrayList<>(input);
      temp.remove(i);
      if (validateSafe(temp)) return true;
    }
    return false;
  }

  public static int countSafe(List<List<Integer>> inputList){
    int count = 0;
    for (List<Integer> list: inputList){
      if (validateSafe(list)) count++;
    }
    return count;
  }

  public static int countSafeAfterRemBadLevel(List<List<Integer>> inputList){
    int count = 0;
    for (List<Integer> list: inputList){
      if (isSafeAfterRemBadLevel(list)) count++;
    }
    return count;
  }

  public static void main(String[] args) {
    List<List<Integer>> inputList = new ArrayList<>();
    try (BufferedReader br = new BufferedReader(new FileReader("./input.txt"))){
      String input;
      while ((input = br.readLine()) != null){
        List<Integer> inputs = parseInput(input);
        inputList.add(inputs);
      }
    } catch(Exception e){
      System.out.println(e.getMessage());
    }
    System.out.println(countSafe(inputList));
    System.out.println(countSafeAfterRemBadLevel(inputList));
  }
}