import numpy as np
import matplotlib.pyplot as plt
import pandas as pd
from sklearn.preprocessing import LabelEncoder, StandardScaler
from sklearn.model_selection import train_test_split
import seaborn as sns

# =============================================================================
# DATA LOADING AND PREPROCESSING
# =============================================================================

# Load the dataset
data = pd.read_csv('/Users/mohsin/MCA/Assignments/AIML/SET-2-Program-2-LDA.csv')

print("Dataset Shape:", data.shape)
print("\nDataset Info:")
print(data.info())
print("\nFirst few rows:")
print(data.head())
print("\nClass distribution:")
print(data['Class'].value_counts())

# =============================================================================
# PROGRAM-1: PERCEPTRON IMPLEMENTATION
# =============================================================================

class Perceptron:
    def __init__(self, learning_rate=0.1, n_iterations=100):
        self.learning_rate = learning_rate
        self.n_iterations = n_iterations
        self.errors = []
        self.weights = None
        self.bias = None
    
    def fit(self, X, y, weight_range=(0.2, 0.8)):
        # Initialize random weights in specified range
        n_features = X.shape[1]
        self.weights = np.random.uniform(weight_range[0], weight_range[1], n_features)
        self.bias = 0
        
        for i in range(self.n_iterations):
            errors = 0
            for idx, x_i in enumerate(X):
                # Forward pass
                linear_output = np.dot(x_i, self.weights) + self.bias
                y_predicted = np.where(linear_output >= 0, 1, 0)
                
                # Update weights
                update = self.learning_rate * (y[idx] - y_predicted)
                self.weights += update * x_i
                self.bias += update
                errors += int(update != 0.0)
            
            self.errors.append(errors)
            if errors == 0: 
                print(f"Perceptron converged at iteration {i+1}")
                break
        
        return self
    
    def predict(self, X):
        linear_output = np.dot(X, self.weights) + self.bias
        return np.where(linear_output >= 0, 1, 0)
    
    def accuracy(self, X, y):
        predictions = self.predict(X)
        return np.mean(predictions == y)

# Prepare data for Perceptron
# Select numerical features for perceptron
numerical_features = ['raisedhands', 'VisITedResources', 'AnnouncementsView', 'Discussion']
X_perceptron = data[numerical_features].values

# Convert class labels to binary (H=1, L=0)
le = LabelEncoder()
y_perceptron = le.fit_transform(data['Class'])

# Normalize features
scaler = StandardScaler()
X_perceptron_scaled = scaler.fit_transform(X_perceptron)

# Split data
X_train, X_test, y_train, y_test = train_test_split(
    X_perceptron_scaled, y_perceptron, test_size=0.3, random_state=42, stratify=y_perceptron
)

print("\n" + "="*50)
print("PERCEPTRON ANALYSIS")
print("="*50)

# Train Perceptron with different weight ranges
print("\n1. Training Perceptron with weight range (0.2, 0.8):")
p1 = Perceptron(learning_rate=0.1, n_iterations=100)
p1.fit(X_train, y_train, (0.2, 0.8))
train_acc1 = p1.accuracy(X_train, y_train)
test_acc1 = p1.accuracy(X_test, y_test)
print(f"Training Accuracy: {train_acc1:.4f}")
print(f"Testing Accuracy: {test_acc1:.4f}")

print("\n2. Training Perceptron with weight range (0.4, 0.7):")
p2 = Perceptron(learning_rate=0.1, n_iterations=100)
p2.fit(X_train, y_train, (0.4, 0.7))
train_acc2 = p2.accuracy(X_train, y_train)
test_acc2 = p2.accuracy(X_test, y_test)
print(f"Training Accuracy: {train_acc2:.4f}")
print(f"Testing Accuracy: {test_acc2:.4f}")

# Plot error functions
plt.figure(figsize=(12, 4))
plt.subplot(1, 2, 1)
plt.plot(p1.errors, label='Weights (0.2-0.8)', marker='o')
plt.plot(p2.errors, label='Weights (0.4-0.7)', marker='s')
plt.xlabel('Iterations')
plt.ylabel('Number of Errors')
plt.title('Perceptron Error Function')
plt.legend()
plt.grid(True)

# 3. NOT gate implementation
print("\n3. NOT Gate Implementation:")
def not_gate(x):
    """NOT gate with weight=-1, bias=0.5"""
    return 1 if (-1 * x + 0.5) >= 0 else 0

print("NOT Gate Truth Table:")
for x in [0, 1]:
    output = not_gate(x)
    print(f"Input: {x}, Output: {output}")

# =============================================================================
# PROGRAM-2: LDA IMPLEMENTATION
# =============================================================================

print("\n" + "="*50)
print("LINEAR DISCRIMINANT ANALYSIS (LDA)")
print("="*50)

# Prepare data for LDA - using numerical features
X_lda = data[numerical_features].values
y_lda = data['Class'].values

# Separate classes
X_H = X_lda[y_lda == 'H']  # High performance
X_L = X_lda[y_lda == 'L']  # Low performance

print(f"\nClass H samples: {len(X_H)}")
print(f"Class L samples: {len(X_L)}")

# Q1: Class means
mean_H = np.mean(X_H, axis=0)
mean_L = np.mean(X_L, axis=0)

print(f"\n1. Class Means:")
print(f"Mean of Class H: {mean_H}")
print(f"Mean of Class L: {mean_L}")

# Q2: Covariance matrices
cov_H = np.cov(X_H.T)
cov_L = np.cov(X_L.T)

print(f"\n2. Covariance Matrices:")
print(f"Covariance of Class H:\n{cov_H}")
print(f"Covariance of Class L:\n{cov_L}")

# Q3: Within-class scatter matrix
S_W = cov_H * (len(X_H) - 1) + cov_L * (len(X_L) - 1)

print(f"\n3. Within-class Scatter Matrix S_W:\n{S_W}")

# Q4: Between-class scatter matrix (Manual Formula Implementation)
print(f"\n4. Between-class Scatter Matrix (Manual Implementation):")

# Calculate overall mean manually
n_samples = len(X_lda)
overall_mean = np.zeros(X_lda.shape[1])
for i in range(X_lda.shape[1]):
    overall_mean[i] = np.sum(X_lda[:, i]) / n_samples

print(f"Overall mean (manual): {overall_mean}")

# Manual calculation of S_B using the formula: S_B = Σ n_i * (μ_i - μ) * (μ_i - μ)^T
n_H = len(X_H)
n_L = len(X_L)

# Calculate mean differences manually
mean_diff_H = mean_H - overall_mean
mean_diff_L = mean_L - overall_mean

print(f"Mean difference H: {mean_diff_H}")
print(f"Mean difference L: {mean_diff_L}")

# Manual outer product calculation for S_B
S_B = np.zeros((X_lda.shape[1], X_lda.shape[1]))

# For class H: n_H * (μ_H - μ) * (μ_H - μ)^T
for i in range(len(mean_diff_H)):
    for j in range(len(mean_diff_H)):
        S_B[i, j] += n_H * mean_diff_H[i] * mean_diff_H[j]

# For class L: n_L * (μ_L - μ) * (μ_L - μ)^T
for i in range(len(mean_diff_L)):
    for j in range(len(mean_diff_L)):
        S_B[i, j] += n_L * mean_diff_L[i] * mean_diff_L[j]

print(f"Between-class Scatter Matrix S_B (manual):\n{S_B}")

# Q5: Manual Matrix Inversion and Multiplication
print(f"\n5. S_W^(-1) * S_B (Manual Implementation):")

def manual_matrix_inverse(matrix):
    """Manual matrix inversion using Gauss-Jordan elimination"""
    n = matrix.shape[0]
    # Create augmented matrix [A|I]
    augmented = np.hstack([matrix.copy(), np.eye(n)])
    
    # Forward elimination
    for i in range(n):
        # Find pivot
        max_row = i
        for k in range(i+1, n):
            if abs(augmented[k, i]) > abs(augmented[max_row, i]):
                max_row = k
        
        # Swap rows
        if max_row != i:
            augmented[[i, max_row]] = augmented[[max_row, i]]
        
        # Make diagonal element 1
        pivot = augmented[i, i]
        if abs(pivot) < 1e-10:
            continue  # Skip if pivot is too small
        
        for j in range(2*n):
            augmented[i, j] /= pivot
        
        # Eliminate column
        for k in range(n):
            if k != i:
                factor = augmented[k, i]
                for j in range(2*n):
                    augmented[k, j] -= factor * augmented[i, j]
    
    return augmented[:, n:]

try:
    # Manual matrix inversion
    S_W_inv = manual_matrix_inverse(S_W)
    print(f"S_W inverse (manual):\n{S_W_inv}")
    
    # Manual matrix multiplication
    SW_inv_SB = np.zeros((S_W.shape[0], S_B.shape[1]))
    for i in range(S_W_inv.shape[0]):
        for j in range(S_B.shape[1]):
            for k in range(S_W_inv.shape[1]):
                SW_inv_SB[i, j] += S_W_inv[i, k] * S_B[k, j]
    
    print(f"S_W^(-1) * S_B (manual):\n{SW_inv_SB}")
    
    # Q6: Manual Eigenvalue Decomposition using Power Iteration Method
    print(f"\n6. Eigenvalues and Eigenvectors (Manual Implementation):")
    
    def power_iteration(matrix, num_iterations=1000, tolerance=1e-10):
        """Find dominant eigenvalue and eigenvector using power iteration"""
        n = matrix.shape[0]
        # Initialize random vector
        v = np.random.rand(n)
        v = v / np.linalg.norm(v)
        
        for i in range(num_iterations):
            # Matrix-vector multiplication
            Av = np.zeros(n)
            for j in range(n):
                for k in range(n):
                    Av[j] += matrix[j, k] * v[k]
            
            # Calculate eigenvalue (Rayleigh quotient)
            eigenvalue = np.dot(v, Av)
            
            # Normalize
            norm = np.sqrt(np.sum(Av**2))
            if norm < tolerance:
                break
            v_new = Av / norm
            
            # Check convergence
            if np.linalg.norm(v_new - v) < tolerance:
                break
            v = v_new
        
        return eigenvalue, v
    
    # Find dominant eigenvalue and eigenvector
    eigenvalue_1, eigenvector_1 = power_iteration(SW_inv_SB)
    
    print(f"Dominant eigenvalue (manual): {eigenvalue_1}")
    print(f"Dominant eigenvector (manual): {eigenvector_1}")
    
    # Use the dominant eigenvector as LDA projection vector
    W_lda = eigenvector_1.reshape(-1, 1)
    print(f"LDA Projection Vector W (manual):\n{W_lda}")
    
    eigenvalues = np.array([eigenvalue_1])  # For compatibility with rest of code
    eigenvectors = eigenvector_1.reshape(-1, 1)
    
except Exception as e:
    print(f"Manual implementation failed: {e}")
    print("Falling back to numpy implementation...")
    # Fallback to numpy
    SW_inv_SB = np.dot(np.linalg.inv(S_W), S_B)
    eigenvalues, eigenvectors = np.linalg.eig(SW_inv_SB)
    
    # Sort eigenvalues and eigenvectors
    sorted_idx = np.argsort(eigenvalues.real)[::-1]
    eigenvalues = eigenvalues.real[sorted_idx]
    eigenvectors = eigenvectors.real[:, sorted_idx]
    
    print(f"\n6. Eigenvalues: {eigenvalues}")
    print(f"Eigenvectors:\n{eigenvectors}")
    
    # Select the eigenvector with largest eigenvalue
    W_lda = eigenvectors[:, 0].reshape(-1, 1)
    print(f"\nLDA Projection Vector W:\n{W_lda}")

# Q7: Transform data (works for both manual and numpy implementations)
X_lda_projected = np.dot(X_lda, W_lda)

print(f"\n7. Projected Data Shape: {X_lda_projected.shape}")
print(f"First 10 projected values: {X_lda_projected[:10].flatten()}")

# Q8: Visualization
plt.subplot(1, 2, 2)
plt.scatter(X_lda_projected[y_lda=='H'], np.zeros(sum(y_lda=='H')), 
           c='red', label='High (H)', alpha=0.7, s=50)
plt.scatter(X_lda_projected[y_lda=='L'], np.zeros(sum(y_lda=='L')), 
           c='blue', label='Low (L)', alpha=0.7, s=50)
plt.xlabel('LDA Projection')
plt.ylabel('Classes')
plt.title('LDA 1D Projection')
plt.legend()
plt.grid(True)

plt.tight_layout()
plt.show()

# =============================================================================
# PROGRAM-3: PCA IMPLEMENTATION
# =============================================================================

print("\n" + "="*50)
print("PRINCIPAL COMPONENT ANALYSIS (PCA)")
print("="*50)

# Use all numerical features for PCA
X_pca = data[numerical_features].values

print(f"Original data shape: {X_pca.shape}")
print(f"Features used: {numerical_features}")

# Q9: Standardization
mean_pca = np.mean(X_pca, axis=0)
std_pca = np.std(X_pca, axis=0, ddof=1)
X_standardized = (X_pca - mean_pca) / std_pca

print(f"\n9. Standardization:")
print(f"Original means: {mean_pca}")
print(f"Original std: {std_pca}")
print(f"Standardized means: {np.mean(X_standardized, axis=0)}")
print(f"Standardized std: {np.std(X_standardized, axis=0, ddof=1)}")

# Q10: Covariance matrix (manual calculation)
mean_centered = X_standardized - np.mean(X_standardized, axis=0)
cov_matrix = np.dot(mean_centered.T, mean_centered) / (len(X_pca) - 1)

print(f"\n10. Covariance Matrix:\n{cov_matrix}")

# Q11: Eigenvalues and eigenvectors
eigenvalues_pca, eigenvectors_pca = np.linalg.eig(cov_matrix)

print(f"\n11. Eigenvalues: {eigenvalues_pca}")
print(f"Eigenvectors:\n{eigenvectors_pca}")

# Q12: Sort by eigenvalues
sorted_idx_pca = np.argsort(eigenvalues_pca)[::-1]
eigenvalues_pca = eigenvalues_pca[sorted_idx_pca]
eigenvectors_pca = eigenvectors_pca[:, sorted_idx_pca]

print(f"\n12. Sorted Eigenvalues: {eigenvalues_pca}")
print(f"Sorted Eigenvectors:\n{eigenvectors_pca}")

# Calculate explained variance ratio
explained_variance_ratio = eigenvalues_pca / np.sum(eigenvalues_pca)
cumulative_variance = np.cumsum(explained_variance_ratio)

print(f"\nExplained Variance Ratio: {explained_variance_ratio}")
print(f"Cumulative Variance: {cumulative_variance}")

# Q13: Select k components
k = 2
selected_eigenvectors = eigenvectors_pca[:, :k]

print(f"\n13. Selected {k} components:")
print(f"Selected eigenvectors:\n{selected_eigenvectors}")
print(f"Variance explained by first {k} components: {cumulative_variance[k-1]:.4f}")

# Q14: Transform data
X_pca_transformed = np.dot(X_standardized, selected_eigenvectors)

print(f"\n14. Transformed data shape: {X_pca_transformed.shape}")
print(f"First 5 transformed samples:\n{X_pca_transformed[:5]}")

# Q15: Visualize PCA results
plt.figure(figsize=(15, 5))

# PCA scatter plot
plt.subplot(1, 3, 1)
colors = ['red' if c == 'H' else 'blue' for c in y_lda]
plt.scatter(X_pca_transformed[:, 0], X_pca_transformed[:, 1], c=colors, alpha=0.7)
plt.xlabel(f'PC1 ({explained_variance_ratio[0]*100:.1f}%)')
plt.ylabel(f'PC2 ({explained_variance_ratio[1]*100:.1f}%)')
plt.title('PCA Result - First 2 Components')
plt.grid(True)

# Explained variance plot
plt.subplot(1, 3, 2)
plt.bar(range(1, len(eigenvalues_pca)+1), explained_variance_ratio, alpha=0.7)
plt.xlabel('Principal Component')
plt.ylabel('Explained Variance Ratio')
plt.title('Explained Variance by Component')
plt.grid(True)

# Cumulative variance plot
plt.subplot(1, 3, 3)
plt.plot(range(1, len(cumulative_variance)+1), cumulative_variance, 'bo-')
plt.xlabel('Number of Components')
plt.ylabel('Cumulative Explained Variance')
plt.title('Cumulative Explained Variance')
plt.grid(True)

plt.tight_layout()
plt.show()

# =============================================================================
# ADDITIONAL ANALYSIS AND INSIGHTS
# =============================================================================

print("\n" + "="*50)
print("ADDITIONAL ANALYSIS AND INSIGHTS")
print("="*50)

# Feature correlation analysis
plt.figure(figsize=(10, 8))
correlation_matrix = data[numerical_features].corr()
sns.heatmap(correlation_matrix, annot=True, cmap='coolwarm', center=0)
plt.title('Feature Correlation Matrix')
plt.show()

# Statistical summary
print("\nStatistical Summary of Numerical Features:")
print(data[numerical_features].describe())

print("\nClass-wise Statistical Summary:")
for class_label in ['H', 'L']:
    print(f"\nClass {class_label}:")
    class_data = data[data['Class'] == class_label][numerical_features]
    print(class_data.describe())

# Performance comparison
print("\n" + "="*50)
print("ALGORITHM PERFORMANCE SUMMARY")
print("="*50)

print(f"\nPerceptron Results:")
print(f"- Weight range (0.2-0.8): Train Acc = {train_acc1:.4f}, Test Acc = {test_acc1:.4f}")
print(f"- Weight range (0.4-0.7): Train Acc = {train_acc2:.4f}, Test Acc = {test_acc2:.4f}")

print(f"\nLDA Results:")
print(f"- Successfully projected {len(X_lda)} samples to 1D")
print(f"- Separation achieved between classes H and L")

print(f"\nPCA Results:")
print(f"- First 2 components explain {cumulative_variance[1]*100:.2f}% of variance")
print(f"- Dimensionality reduced from {X_pca.shape[1]} to {k} features")

print("\n" + "="*50)
print("ANALYSIS COMPLETE")
print("="*50)
