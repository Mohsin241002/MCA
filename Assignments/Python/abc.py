import pandas as pd
import matplotlib.pyplot as plt
from matplotlib.table import Table

# Dataset table content
data = {
    "Variable": [
        "Subscribers", "Visits", "Likes", "Comments", "Rank", "Country", "Categories"
    ],
    "Description": [
        "Total number of subscribers for each channel",
        "Total views accumulated by the channel",
        "Total likes received on the channel's videos",
        "Total comments received on the channel's videos",
        "Rank of the channel based on engagement",
        "Country where the channel is based",
        "Category of content (e.g., Music, Animation)"
    ],
    "Data Type": [
        "Numeric", "Numeric", "Numeric", "Numeric", "Numeric", "Categorical", "Categorical"
    ]
}

# Create a DataFrame
df = pd.DataFrame(data)

# Create the table image
fig, ax = plt.subplots(figsize=(10, 4))
ax.axis('off')

# Add the table to the plot
table = Table(ax, bbox=[0, 0, 1, 1])
n_rows, n_cols = df.shape

# Add table header with bold font
for col_idx, col_name in enumerate(df.columns):
    table.add_cell(0, col_idx, width=1/n_cols, height=0.1, text=col_name,
                   loc='center', facecolor='lightgray', fontproperties={'weight': 'bold'})

# Add table rows
for row_idx, row in df.iterrows():
    for col_idx, value in enumerate(row):
        table.add_cell(row_idx + 1, col_idx, width=1/n_cols, height=0.1, text=value,
                       loc='center')

# Adjust table appearance
for row_idx in range(n_rows + 1):
    for col_idx in range(n_cols):
        cell = table[row_idx, col_idx]
        cell.set_edgecolor('black')
        cell.set_facecolor('white')

ax.add_table(table)

# Save the image
plt.savefig("dataset_table_bold.png", dpi=300, bbox_inches='tight')
plt.show()
