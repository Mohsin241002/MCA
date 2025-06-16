# Import libraries
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
from sklearn.preprocessing import StandardScaler
from sklearn.cluster import AgglomerativeClustering
from scipy.cluster.hierarchy import dendrogram, linkage
from sklearn.metrics import silhouette_score

# 1. Loading the dataset
df = pd.read_csv('Absenteeism_at_work_with_reasons.csv', delimiter=';')

# 2. Visualize the data
# Plotting two graphs
plt.figure()
ax=sns.distplot(df.iloc[:,0], kde=False)
ax.set_title('Distribution of first variable')
plt.show()

plt.figure()
ax=sns.scatterplot(x=df.iloc[:,0], y=df.iloc[:,1])
ax.set_title('Scatter plot of first two variables')
plt.show()

# 3. Identify missing values
print("Missing Values:")
print(df.isnull().sum())

# 4. Select relevant features
# Here we select numerical columns
numerical_features = df.select_dtypes(include=['number']).columns
X = df[numerical_features]

# 5. Normalize the data
scaler = StandardScaler()
X_scaled = scaler.fit_transform(X)

# 6. Dendrogram
Z = linkage(X_scaled, method='ward')
plt.figure()
dendrogram(Z)
plt.title('Dendrogram')
plt.show()

# 7. Implement agglomerative clustering
linkage_methods = ['single', 'complete', 'average', 'ward']
labels = {}
for method in linkage_methods:
    clustering = AgglomerativeClustering(n_clusters=5, linkage=method)
    clusters = clustering.fit_predict(X_scaled)
    silhouette = silhouette_score(X_scaled, clusters)
    print(f"Silhouette score with {method} linkage: {silhouette}")
    labels[method] = clusters

# 8. Visualize clusters
# Plot first two principal components
from sklearn.decomposition import PCA
pca = PCA(n_components=2)
X_pca = pca.fit_transform(X_scaled)

plt.figure()
for method, clusters in labels.items():
    plt.scatter(X_pca[:,0], X_pca[:,1], c=clusters, label=method)

plt.legend()
plt.title('Agglomerative Clustering')
plt.xlabel('PCA 1')
plt.ylabel('PCA 2')
plt.show()
