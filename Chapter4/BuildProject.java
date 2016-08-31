import myGraph.*;
import java.util.*;
public class BuildProject {
	public static GraphNode[] buildGraph(String[] pList, Map dep) {
		GraphNode[] projList = new GraphNode[pList.length];
		int count = 0;
		for(String p:pList) {
			GraphNode n1 = new GraphNode(p);
			if(count < projList.length) {
				projList[count] = n1;
				count++;
			}
			if(dep.containsKey(p)) {
				GraphNode nb = new GraphNode((String)dep.get(p));
				if(!Arrays.asList(projList).contains(nb)) {
					if(count < projList.length) {
						projList[count] = nb;
						count++;
					}
				}
				n1.setNeighbor(nb);
			}
		}
		return projList;
	}

	public static void buildOrder(GraphNode[] n, GraphNode[] order, Map visit, int nCount, int oCount) {
		GraphNode curr = n[nCount];
		if((nCount > n.length || oCount > order.length) && visit.size() < n.length) {
			System.out.println("No order can be formed");
		}

		else if(curr != null && !checkVisited(curr, visit)) {
			if(!curr.hasNeighbor()) {
				visit.put(curr.name, true);
				order[oCount] = curr;
				oCount++;
				nCount++;	
			}
			else {
				Boolean visitNeighbor = true;

				for(int i = 0; i<curr.getNeighborsCount(); i++) {
					visitNeighbor = checkVisited(curr.neighbors[i], visit);
					if(!visitNeighbor) {
						break;
					}
					else {
						visit.put(curr.neighbors[i].name, true);
						order[oCount] = curr.neighbors[i];
						oCount++;
						nCount++;	
					}						
				}

				if(!visitNeighbor) {
					nCount++;
					buildOrder(n, order, visit, nCount, oCount);
				} 
			}			
		}
	}

	public static Boolean checkVisited(GraphNode n, Map visit) {
		if(visit.containsKey(n.name)) {
			return true;
		}
		else {
			return false;
		}
	} 

	public static void main(String[] args) {
		String[] projects = new String[6];
		projects[0] = "a";
		projects[1] = "b";
		projects[2] = "c";
		projects[3] = "d";
		projects[4] = "e";
		projects[5] = "f";
		Map<String, String> dependency = new HashMap<String, String>();
		dependency.put("a", "f");
		dependency.put("b", "f");
		dependency.put("c", "d");
		dependency.put("d", "a");
		dependency.put("d", "b");

		GraphNode[] nList = buildGraph(projects, dependency);
		GraphNode[] ord = new GraphNode[nList.length];	
		Map<String, Boolean> visitData = new HashMap<String, Boolean>();
		buildOrder(nList, ord, visitData, 0, 0);
		for(int i = 0; i<ord.length; i++) {
			System.out.println(ord[i].name);
		}
	}
}












