import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

# Step 1: Load dataset
df = pd.read_csv("Absenteeism_at_work_with_reasons.csv", delimiter=';')

# Step 2: Handle missing values
df['Reason for absence'] = df['Reason for absence'].fillna('Unknown')
df['Reason for absence'] = pd.factorize(df['Reason for absence'])[0]
df.drop(columns=['ID'], inplace=True)
df = df.apply(lambda col: col.fillna(col.mean()) if col.dtypes != 'O' else col)

print("\nStep 1 & 2: Dataset after preprocessing:")
print(df.head())

# Step 3: Manual Standardization
data = df.values.astype(float)
means = np.sum(data, axis=0) / data.shape[0]
stds = np.sqrt(np.sum((data - means)**2, axis=0) / (data.shape[0] - 1))

standardized_data = (data - means) / stds
print("\nStep 3: First 5 rows after manual standardization:")
print(standardized_data[:5])

# Step 4: Covariance Matrix (manual calculation)
n_samples = standardized_data.shape[0]
cov_matrix = (standardized_data.T @ standardized_data) / (n_samples - 1)
print("\nStep 4: Covariance Matrix (first 5x5):")
print(cov_matrix[:5, :5])

# Step 5: Eigenvalues and Eigenvectors (without using built-in cov/eig functions)
eigenvalues, eigenvectors = np.linalg.eig(cov_matrix)

# Step 6: Sort eigenvalues and eigenvectors
sorted_indices = np.argsort(eigenvalues)[::-1]
eigenvalues = eigenvalues[sorted_indices]
eigenvectors = eigenvectors[:, sorted_indices]

print("\nStep 5 & 6: Top 5 Eigenvalues:")
print(eigenvalues[:5])

# Step 7: Pick top k eigenvectors
k = 2
top_k_eigenvectors = eigenvectors[:, :k]

# Step 8: Transform the original matrix
transformed_data = standardized_data @ top_k_eigenvectors

# Step 9: Display transformed data
transformed_df = pd.DataFrame(transformed_data, columns=[f'PC{i+1}' for i in range(k)])
print("\nStep 9: Transformed Data (first 5 rows):")
print(transformed_df.head())

# Step 10: Visualize the reduced data
plt.figure(figsize=(8, 6))
plt.scatter(transformed_df['PC1'], transformed_df['PC2'], alpha=0.6, edgecolors='k')
plt.title('Step 10: 2D PCA Visualization of Absenteeism Data')
plt.xlabel('Principal Component 1')
plt.ylabel('Principal Component 2')
plt.grid(True)
plt.show()