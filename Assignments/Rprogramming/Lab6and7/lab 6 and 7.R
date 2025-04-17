library(ggplot2)  
library(dplyr)    
file_path <- "D:/MCA/Assignments/Rprogramming/Lab6and7/training_productivity_anova.csv"  
data <- read.csv(file_path)

# columns to factors
data$Department <- as.factor(data$Department)
data$Training_Method <- as.factor(data$Training_Method)

# One-Way ANOVA - Effect of Training_Method on Productivity
one_way_anova <- aov(Productivity ~ Training_Method, data = data)
summary(one_way_anova)

# Two-Way ANOVA - Effect of Training_Method and Department on Productivity
two_way_anova <- aov(Productivity ~ Training_Method * Department, data = data)
summary(two_way_anova)


# Boxplot of Productivity by Training Method
ggplot(data, aes(x = Training_Method, y = Productivity, fill = Training_Method)) +
  geom_boxplot() +
  labs(title = "Boxplot of Productivity by Training Method", x = "Training Method", y = "Productivity") +
  theme_minimal()

# Boxplot of Productivity by Department
ggplot(data, aes(x = Department, y = Productivity, fill = Department)) +
  geom_boxplot() +
  labs(title = "Boxplot of Productivity by Department", x = "Department", y = "Productivity") +
  theme_minimal()

# Boxplot of Productivity by Training Method and Department
ggplot(data, aes(x = Training_Method, y = Productivity, fill = Department)) +
  geom_boxplot() +
  facet_wrap(~ Department) +
  labs(title = "Boxplot of Productivity by Training Method and Department", x = "Training Method", y = "Productivity") +
  theme_minimal()

# Density plot of Productivity
ggplot(data, aes(x = Productivity, fill = Training_Method)) +
  geom_density(alpha = 0.6) +
  labs(title = "Density Plot of Productivity by Training Method", x = "Productivity", y = "Density") +
  theme_minimal()

# Density plot of Productivity by Department
ggplot(data, aes(x = Productivity, fill = Department)) +
  geom_density(alpha = 0.6) +
  labs(title = "Density Plot of Productivity by Department", x = "Productivity", y = "Density") +
  theme_minimal()

# Printing Summaries
cat("\nOne-Way ANOVA Results:\n")
print(summary(one_way_anova))

cat("\nTwo-Way ANOVA Results:\n")
print(summary(two_way_anova))
