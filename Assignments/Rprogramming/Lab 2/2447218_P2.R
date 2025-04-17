# Step 1: Create vectors for student names and marks for three subjects
# Three students: Snehil, Rohan, and Mohsin
students <- c("Snehil", "Rohan", "Mohsin")
math_marks <- c(85, 78, 92)
science_marks <- c(88, 74, 90)
english_marks <- c(90, 80, 85)

# Step 2: Create a matrix where rows represent students and columns represent subjects
student_matrix <- matrix(c(math_marks, science_marks, english_marks),
                         nrow = 3, ncol = 3, byrow = TRUE)

# Assign row and column names to the matrix
colnames(student_matrix) <- c("Math", "Science", "English")
rownames(student_matrix) <- students

# Print the matrix to show that rows are students and columns are subjects
cat("Original Matrix (Rows = Students, Columns = Subjects):\n")
print(student_matrix)

# Step 3: Transpose the matrix so that subjects become rows and students become columns
transposed_student_matrix <- t(student_matrix)

# Print the transposed matrix to show that rows are now subjects and columns are students
cat("\nTransposed Matrix (Rows = Subjects, Columns = Students):\n")
print(transposed_student_matrix)

# Step 4: Create midterm and endterm dataframes with students' marks
midterm <- data.frame(
  Student = c("Snehil", "Rohan", "Mohsin"),
  Math = c(85, 78, 92),
  Science = c(88, 74, 90),
  English = c(90, 80, 85)
)

endterm <- data.frame(
  Student = c("Snehil", "Rohan", "Mohsin"),
  Math = c(80, 82, 89),
  Science = c(85, 79, 87),
  English = c(92, 84, 88)
)

# Print midterm and endterm dataframes
cat("\nMidterm Marks:\n")
print(midterm)
cat("\nEndterm Marks:\n")
print(endterm)

# Step 5: Merge midterm and endterm by student names
merged_df <- merge(midterm, endterm, by = "Student", suffixes = c("_midterm", "_endterm"))

# Print the merged dataframe
cat("\nMerged Dataframe (Midterm and Endterm Marks):\n")
print(merged_df)

# Step 6: Compute total and average marks for each student
merged_df$Total_Math <- merged_df$Math_midterm + merged_df$Math_endterm
merged_df$Total_Science <- merged_df$Science_midterm + merged_df$Science_endterm
merged_df$Total_English <- merged_df$English_midterm + merged_df$English_endterm

# Compute the average marks for each subject
merged_df$Average_Math <- merged_df$Total_Math / 2
merged_df$Average_Science <- merged_df$Total_Science / 2
merged_df$Average_English <- merged_df$Total_English / 2

# Print the dataframe with total and average marks
cat("\nDataframe with Total and Average Marks:\n")
print(merged_df)

# Step 7: Melt the merged dataframe into long format
# Reshape the dataframe so that each row contains one subject-term combination
library(reshape2)


melted_df <- melt(merged_df, id.vars = "Student", 
                  measure.vars = c("Math_midterm", "Science_midterm", "English_midterm", 
                                   "Math_endterm", "Science_endterm", "English_endterm"),
                  variable.name = "Subject_Term", value.name = "Marks")

# Print the melted dataframe
cat("\nMelted Dataframe (Long Format):\n")
print(melted_df)

# Step 8: Cast the melted data back to wide format for further analysis
# Reshape the melted dataframe back to a wide format, with each subject-term combination in its own column
wide_df <- dcast(melted_df, Student ~ Subject_Term, value.var = "Marks")

# Print the wide dataframe
cat("\nWide Dataframe (Re-casted from Long Format):\n")
print(wide_df)
