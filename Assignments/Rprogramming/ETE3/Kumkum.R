# Problem Statement:
# The YouTube dataset contains key metrics such as subscribers, visits, likes, comments, ranks, categories, and countries.
# The goal is to analyze relationships, distributions, and engagement trends across these variables, 
# focusing on the top 5 countries and categories. This analysis uses descriptive statistics, visualizations, 
# and statistical tests (ANOVA) to gain insights into YouTube channel performance and audience behavior.

# Loading necessary libraries
library(ggplot2)
library(dplyr)
library(scales)
library(treemap)

# Overview of the Study
# This report analyzes data from YouTube channels, focusing on various variables including the number of subscribers, visits, likes, comments, and ranks. 
# The aim is to explore relationships between these variables, particularly focusing on the distribution of YouTube channel ranks by country and category.
# We also investigate the number of YouTubers from different countries.

# Dataset Description (Metadata)
# The dataset used in this report contains information on YouTube channels. The data includes variables such as:
# - Rank: The rank of the YouTube channel based on some criteria.
# - Country: The country where the YouTube channel is based.
# - Categories: The categories under which the channels fall (e.g., Music, Animation, etc.).
# - Subscribers: The number of subscribers each channel has.
# - Visits: The number of visits each channel has accumulated.
# - Likes: The total number of likes each channel has received.
# - Comments: The total number of comments each channel has received.

# Data Preprocessing
# Load the dataset
data <- read.csv("D:/MCA/Assignments/Rprogramming/ETE3/Youtube.csv")

# Data Preprocessing
data$Likes <- as.numeric(data$Likes)
data$Comments <- as.numeric(data$Comments)
data$Rank <- as.numeric(as.character(data$Rank))

# Remove "Unknown" countries and "-" categories
data <- data %>%
  filter(Country != "Unknown", Categories != "-")

# Filter the Top 5 Categories and Top 5 Countries
top_categories <- data %>%
  count(Categories) %>%
  top_n(5, n) %>%
  pull(Categories)

top_countries <- data %>%
  count(Country) %>%
  top_n(5, n) %>%
  pull(Country)

data_filtered <- data %>%
  filter(Categories %in% top_categories & Country %in% top_countries)

# Descriptive Analysis
# Calculate Descriptive Statistics for Subscribers, Visits, Likes, and Comments
descriptive_stats <- data.frame(
  Metric = c("Mean", "Median", "Mode", "Variance", "Standard Deviation", "Range"),
  Subscribers = c(
    mean(data_filtered$Suscribers, na.rm = TRUE),
    median(data_filtered$Suscribers, na.rm = TRUE),
    as.numeric(names(sort(table(data_filtered$Suscribers), decreasing = TRUE))[1]),
    var(data_filtered$Suscribers, na.rm = TRUE),
    sd(data_filtered$Suscribers, na.rm = TRUE),
    diff(range(data_filtered$Suscribers, na.rm = TRUE))
  ),
  Visits = c(
    mean(data_filtered$Visits, na.rm = TRUE),
    median(data_filtered$Visits, na.rm = TRUE),
    as.numeric(names(sort(table(data_filtered$Visits), decreasing = TRUE))[1]),
    var(data_filtered$Visits, na.rm = TRUE),
    sd(data_filtered$Visits, na.rm = TRUE),
    diff(range(data_filtered$Visits, na.rm = TRUE))
  ),
  Likes = c(
    mean(data_filtered$Likes, na.rm = TRUE),
    median(data_filtered$Likes, na.rm = TRUE),
    as.numeric(names(sort(table(data_filtered$Likes), decreasing = TRUE))[1]),
    var(data_filtered$Likes, na.rm = TRUE),
    sd(data_filtered$Likes, na.rm = TRUE),
    diff(range(data_filtered$Likes, na.rm = TRUE))
  ),
  Comments = c(
    mean(data_filtered$Comments, na.rm = TRUE),
    median(data_filtered$Comments, na.rm = TRUE),
    as.numeric(names(sort(table(data_filtered$Comments), decreasing = TRUE))[1]),
    var(data_filtered$Comments, na.rm = TRUE),
    sd(data_filtered$Comments, na.rm = TRUE),
    diff(range(data_filtered$Comments, na.rm = TRUE))
  )
)

print("Descriptive Statistics:")
print(descriptive_stats)

# **Interpretation of Descriptive Statistics**:
# - The mean and median for subscribers, visits, likes, and comments indicate central tendencies.
# - Variances and standard deviations highlight variability in the data, showing significant variations in engagement metrics.
# - Modes provide insight into the most frequently occurring values for each metric.

# Data Visualizations & Exploration

# 6.1: Bar Plot of Rank Distribution by Categories (Using only top 5 categories)
ggplot(data_filtered, aes(x = Categories, y = Rank, fill = Categories)) + 
  geom_bar(stat = "identity", show.legend = FALSE) + 
  labs(title = "Rank Distribution by Categories (Top 5)", x = "Categories", y = "Rank") + 
  theme(axis.text.x = element_text(angle = 90, hjust = 1)) +
  scale_y_continuous(labels = scales::comma) + 
  theme_minimal()

# **Interpretation**:
# The "Music and Dance" category has the highest rank distribution, indicating its popularity. 
# Categories like "Animation" and "Video Games" also show significant presence, while others like "Movies" are comparatively less prominent in rank distribution.

# 6.2: Pie Chart for Categories Distribution (Using only top 5 categories)
category_counts <- table(data_filtered$Categories)
ggplot(as.data.frame(category_counts), aes(x = "", y = Freq, fill = Var1)) + 
  geom_bar(stat = "identity", width = 1) + 
  coord_polar("y") + 
  labs(title = "Distribution of Categories (Top 5)")

# **Interpretation**:
# The pie chart confirms that "Music and Dance" holds the largest share in the top 5 categories, 
# followed by categories like "Animation" and "Movies," indicating their significant presence on the platform.

# 6.3: Boxplot of Rank by Categories (Using only top 5 categories)
ggplot(data_filtered, aes(x = Categories, y = Rank)) + 
  geom_boxplot(fill = "cyan") + 
  labs(title = "Boxplot of Rank by Categories (Top 5)", x = "Categories", y = "Rank")

# **Interpretation**:
# The boxplot shows a higher spread for the "Music and Dance" category, which suggests a larger variation in ranks. 
# Categories like "Movies" and "Animation" have more compact distributions, indicating more consistency in their ranks.

# 6.4: One-Way ANOVA of Rank by Country (Using only top 5 countries)
anova_result <- aov(Rank ~ Country, data = data_filtered)
summary(anova_result)

# **Interpretation**:
# The ANOVA results show significant differences in ranks across countries, suggesting that countries have varying levels of YouTube channel popularity and rank distributions.

# 6.5: Violin Plot of Rank by Country (Using only top 5 countries)
ggplot(data_filtered, aes(x = Country, y = Rank)) + 
  geom_violin() + 
  labs(title = "Violin Plot of Rank by Country (Top 5)", x = "Country", y = "Rank")

# **Interpretation**:
# The violin plot shows that countries like the USA and India have wider distributions of ranks, indicating greater variation in ranks, whereas other countries such as Indonesia have more concentrated rank distributions.

# 6.6: Treemap for number of YouTubers in each Country
treemap_data <- data_filtered %>%
  count(Country) %>%
  rename(YouTubers = n)

treemap(treemap_data, 
        index = "Country", 
        vSize = "YouTubers", 
        vColor = "YouTubers", 
        draw = TRUE, 
        palette = "Set3", 
        title = "Treemap of YouTubers by Country")

# **Interpretation**:
# The treemap clearly shows that the USA has the largest number of YouTubers, followed by countries like India and Brasil. 
# This highlights the dominance of these countries in terms of the number of YouTube content creators.

# Result Summary
# - The **Bar Plot** of rank distribution by categories shows that "Music and Dance" channels dominate the rank distribution.
# - The **Pie Chart** reaffirms that "Music and Dance" holds the largest share in the top 5 categories, followed by "Animation."
# - The **Boxplot** indicates that the ranks within the "Music and Dance" category have a wider spread, meaning more variation in ranks, compared to other categories.
# - The **ANOVA test** indicates that there are significant differences in YouTube channel ranks based on country.
# - The **Violin Plot** of ranks by country shows varying distributions, with the USA and India having more spread in ranks, suggesting more diversity in YouTube channel ranks across countries.
# - The **Treemap** highlights the distribution of YouTubers by country, showing that the USA has the highest number of YouTubers, followed by India and Brasil.

# Conclusion and Insights
# Based on the visualizations and statistical tests, it is clear that YouTube categories like "Music and Dance" dominate in terms of rank and viewership.
# The number of YouTubers in countries like the USA and India shows a strong global presence. The analysis suggests that there are significant variations in rank distributions across countries.
