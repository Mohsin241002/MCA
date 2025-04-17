library(dplyr)
library(ggplot2)

data <- read.csv("D:/MCA/Assignments/Rprogramming/lab 3 and 4/Dataset1_Mortality.csv")

# Part 3: Descriptive Statistics

# Part A: Number of variables and observations
num_variables <- ncol(data)
num_observations <- nrow(data)
cat("Variables:", num_variables, "Observations:", num_observations, "\n")

# Part B: Extract variable 2 and 3 as vectors
vector2 <- data[[2]]
vector3 <- data[[3]]
cat("Vector for Variable 2:\n", vector2, "\n")
cat("Vector for Variable 3:\n", vector3, "\n")

# Part C: Count the different blood groups in the dataset
num_blood_groups <- length(unique(data$BLOOD))
cat("Different Blood Groups:", num_blood_groups, "\n")

# Part D: List unique SMOKE categories
unique_smoke <- unique(data$SMOKE)
cat("Unique SMOKE Categories:\n")
print(unique_smoke)

# Part E: Count individuals with CHOL levels above 300
# Convert CHOL to numeric in case of any non-numeric values
data$CHOL <- as.numeric(as.character(data$CHOL))

# Check for any NA values after conversion
if (any(is.na(data$CHOL))) {
  cat("Warning: Some CHOL values could not be converted to numeric and are set to NA.\n")
}

high_CHOL <- sum(data$CHOL > 300, na.rm = TRUE)
cat("Individuals with CHOL > 300:", high_CHOL, "\n")

# Part F: Mean HEIGHT for individuals where mortality is 'alive'
mean_HEIGHT_alive <- mean(data$HEIGHT[data$MORT == "alive"], na.rm = TRUE)
cat("Mean HEIGHT for Mortality Alive:", mean_HEIGHT_alive, "\n")

# Part G: AGE of the tallest person with O Blood Group
tallest_o <- data %>% filter(BLOOD == "O") %>% arrange(desc(HEIGHT)) %>% slice(1) %>% select(AGE)
cat("AGE of Tallest O Blood Group Person:\n")
print(tallest_o)

# Part H: Count nonsmokers who are alive and below 40 years old
nonsmokers_alive_below_40 <- sum(data$SMOKE == "no" & data$MORT == "alive" & data$AGE < 40, na.rm = TRUE)
cat("Nonsmokers Alive Below 40:", nonsmokers_alive_below_40, "\n")

# Part 4: Data Visualization

# Part A: Single-variable plots
# Bar plot of BloodGroup
barplot(table(data$BLOOD), main="Blood Group Distribution", col="skyblue")

# Pie chart of BloodGroup
pie(table(data$BLOOD), main="Blood Group Distribution")

# Box plot of AGE
boxplot(data$AGE, main="Box Plot of AGE", col="lightgreen")

# Histogram of CHOL
hist(data$CHOL, main="Histogram of CHOL", col="lightblue", xlab="CHOL Levels")

# Part B: Two-variable plots
# Bar plot for Mortality and SMOKE
with(data, barplot(table(MORT, SMOKE), beside=TRUE, col=c("orange", "blue"), main="Mortality vs Smoke Status"))

# Scatter plot for AGE and CHOL
plot(data$AGE, data$CHOL, main="Scatter Plot of AGE vs CHOL", xlab="AGE", ylab="CHOL", col="purple")

# Box plot for AGE by BloodGroup
boxplot(data$AGE ~ data$BLOOD, main="Box Plot of AGE by Blood Group", col="salmon")

# Part C: Multivariable 2D Bar Plot with ggplot2

# Summarize data for the 2D bar plot
mort_smoke_blood_summary <- data %>%
  group_by(MORT, SMOKE, BLOOD) %>%
  summarise(Count = n()) %>%
  ungroup()

# Create a 2D bar plot
ggplot(mort_smoke_blood_summary, aes(x = MORT, y = Count, fill = BLOOD)) +
  geom_bar(stat = "identity", position = "dodge") +
  facet_wrap(~ SMOKE) +
  labs(title = "Counts of MORT by SMOKE and BLOOD Group",
       x = "Mortality Status",
       y = "Count") +
  theme_minimal() +
  theme(legend.position = "bottom")
