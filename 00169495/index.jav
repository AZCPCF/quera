import java.util.HashMap;
import java.util.HashSet;

public class HashCollisionChecker {

    public static <T> int countOfUniqueHashCodes(HashSet<T> set) {
        HashSet<Integer> uniqueHashes = new HashSet<>();
        for (T item : set) {
            uniqueHashes.add(item.hashCode());
        }
        return uniqueHashes.size();
    }

    public static <K, V> int countOfUniqueHashCodes(HashMap<K, V> map) {
        HashSet<Integer> uniqueHashes = new HashSet<>();
        for (K key : map.keySet()) {
            uniqueHashes.add(key.hashCode());
        }
        return uniqueHashes.size();
    }
}
