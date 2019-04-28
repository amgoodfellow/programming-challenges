public class MinChange {
  private int[] denominations;

  public MinChange(int[] denoms) {
    denominations = denoms;
  }

  private int dynamicSolution(int request, int[] memo) {
    int minValue = request;

    if (memo[request] != 0) {
      return memo[request];
    }

    for (int x : this.denominations) {
      if (x > request) {
        continue;
      } else if (x == request) {
        memo[request] = 1;
        return 1;
      }

      int temp = 1 + dynamicSolution(request - x, memo);

      if (temp < minValue) {
        minValue = temp;
        memo[request] = minValue;
      }

    }

    return minValue;
  }

  public static void main

  public static void main(String[] args) {
    int[] denoms = { 50, 20, 1 };
    int request = 63;
    MinChange changer = new MinChange(denoms);

    System.out.println(changer.dynamicSolution(request, new int[request + 1]));

  }

}