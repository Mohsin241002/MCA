import pandas as pd
from scipy.stats import f_oneway, ttest_1samp
import statsmodels.api as sm
from statsmodels.formula.api import ols

# Load your dataset
file_path = "D:/MCA/Assignments/Rprogramming/ETE3/Youtube.csv"
data = pd.read_csv(file_path)

# Preprocessing
data = data.dropna(subset=["Suscribers", "Likes", "Comments", "Rank", "Country", "Categories"])
data = data[data["Country"] != "Unknown"]
data = data[data["Categories"] != "-"]

# One-Way ANOVA: Rank by Country
countries = data.groupby("Country")["Rank"].apply(list)
anova_oneway_result = f_oneway(*countries)

# Two-Way ANOVA: Likes by Country and Categories
data['Country'] = data['Country'].astype('category')
data['Categories'] = data['Categories'].astype('category')
model = ols('Likes ~ C(Country) + C(Categories) + C(Country):C(Categories)', data=data).fit()
anova_twoway_result = sm.stats.anova_lm(model, typ=2)

# T-Test: Subscribers vs Hypothetical Mean (e.g., 1M subscribers)
t_test_result = ttest_1samp(data["Suscribers"], 1_000_000, nan_policy='omit')

# Print Results in Table Format
print("One-Way ANOVA Results:")
print(f"F-statistic: {anova_oneway_result.statistic:.2f}")
print(f"P-value: {anova_oneway_result.pvalue:.4f}")

print("\nTwo-Way ANOVA Results:")
anova_twoway_df = pd.DataFrame(anova_twoway_result)
print(anova_twoway_df)

print("\nT-Test Results:")
print(f"T-statistic: {t_test_result.statistic:.2f}")
print(f"P-value: {t_test_result.pvalue:.4f}")
