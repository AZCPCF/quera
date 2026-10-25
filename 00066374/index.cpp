#include <bits/stdc++.h>
using namespace std;

const int NEG = -1000000000;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n, m;
    cin >> n >> m;

    vector<string> s(n);
    vector<int> len(n);

    for (int i = 0; i < n; i++) {
        cin >> s[i];
        len[i] = (int)s[i].size();
    }

    int N = 1 << m;

    vector<int> pc(N, 0);

    for (int mask = 1; mask < N; mask++) {
        pc[mask] = pc[mask >> 1] + (mask & 1);
    }

    // Masks grouped by their length and number of set bits.
    vector<vector<vector<int>>> byLen(
        m + 1, vector<vector<int>>(m + 1)
    );

    for (int r = 0; r <= m; r++) {
        for (int mask = 0; mask < (1 << r); mask++) {
            byLen[r][pc[mask]].push_back(mask);
        }
    }

    auto horizontal = [&](int mask, const string &str) {
        int pos = 0;
        int score = 0;

        for (int j = 0; j < m; j++) {
            if (mask & (1 << j)) {
                if (j > 0 &&
                    (mask & (1 << (j - 1))) &&
                    pos > 0 &&
                    str[pos - 1] == str[pos]) {
                    score++;
                }

                pos++;
            }
        }

        return score;
    };

    vector<int> dp(N, NEG), ndp(N, NEG);

    // First row.
    for (int mask = 0; mask < N; mask++) {
        if (pc[mask] == len[0]) {
            dp[mask] = horizontal(mask, s[0]);
        }
    }

    // Process the remaining rows.
    for (int i = 1; i < n; i++) {
        int lp = len[i - 1];
        int lc = len[i];

        vector<int> A((size_t)N * (lc + 1), NEG);
        vector<int> B((size_t)N * (lc + 1), NEG);

        // Base case: all columns have been processed.
        for (int mask = 0; mask < N; mask++) {
            A[(size_t)mask * (lc + 1) + lc] = dp[mask];
        }

        // Process columns from right to left.
        for (int j = m - 1; j >= 0; j--) {
            int r = m - j;
            int suffixWidth = 1 << (r - 1);

            int loq = max(0, lc - r);
            int hiq = min(j, lc);

            for (int X = 0; X < (1 << j); X++) {
                int p = pc[X];

                // The previous row must be completable.
                if (p > lp || p + r < lp) {
                    continue;
                }

                bool can0 = (p + r - 1 >= lp);
                bool can1 = (p < lp);

                for (int q = loq; q <= hiq; q++) {
                    int k = lc - q;

                    for (int z : byLen[r][k]) {
                        int b = z & 1;
                        int z2 = z >> 1;
                        int q2 = q + b;

                        int best = NEG;

                        // Previous row leaves this column empty.
                        if (can0) {
                            size_t id =
                                ((size_t)X * (lc + 1) + q2)
                                * suffixWidth + z2;

                            best = A[id];
                        }

                        // Previous row plants a flower here.
                        if (can1) {
                            int X1 = X | (1 << j);

                            size_t id =
                                ((size_t)X1 * (lc + 1) + q2)
                                * suffixWidth + z2;

                            int v = A[id];

                            // Both rows have a flower in this column.
                            if (b &&
                                q < lc &&
                                s[i - 1][p] == s[i][q]) {
                                v++;
                            }

                            best = max(best, v);
                        }

                        size_t out =
                            ((size_t)X * (lc + 1) + q)
                            * (1 << r) + z;

                        B[out] = best;
                    }
                }
            }

            A.swap(B);
        }

        // Add horizontal happiness of the current row.
        fill(ndp.begin(), ndp.end(), NEG);

        for (int mask = 0; mask < N; mask++) {
            if (pc[mask] == lc) {
                ndp[mask] =
                    A[mask] + horizontal(mask, s[i]);
            }
        }

        dp.swap(ndp);
    }

    int ans = 0;

    for (int value : dp) {
        ans = max(ans, value);
    }

    cout << ans << '\n';

    return 0;
}