 import java.util.*;


public class MatrixRotateBF {
		
	private static int[][] rotateMatrix(int[][] inpMat) {
	    int[][] oupMat = new int[3][3];
	    for(int i=0; i<3; i++) {
	        for(int j=0; j<3; j++) {
	            oupMat[j][2-i] = inpMat[i][j];
	        }
	    }
	    return oupMat;
	
	}
	
	
	public static void main(String args[]){
		int[][] mat = new int[3][3];
		
		for(int i = 0; i<3; i++) {
		    for(int j = 0; j<3; j++) {
		        mat[i][j] = i+1;
		    }
		}
		
		for(int i = 0; i<3; i++) {
		    for(int j = 0; j<3; j++) {
		       System.out.print(mat[i][j]);
		    }
		}
		
		int[][] oMat = rotateMatrix(mat);
		System.out.println();
		for(int i = 0; i<3; i++) {
		    for(int j = 0; j<3; j++) {
		       System.out.print(oMat[i][j]);
		    }
		}
	}
}
