{
 "cells": [
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "# Introduction to Panel Data\n",
    "\n",
    "Panel data, also known as longitudinal data, is a dataset that contains observations over time for multiple subjects. It combines both cross-sectional and time-series data. This notebook will help you understand the basics of panel data using Python's Pandas library."
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "## 1. What is Panel Data?\n",
    "\n",
    "- **Cross-sectional data**: Data collected at a single point in time.\n",
    "- **Time series data**: Data collected over time for a single subject.\n",
    "- **Panel data**: Data collected over time for multiple subjects.\n",
    "\n",
    "### Example:\n",
    "Imagine you have data about students' test scores across multiple years. Each student is observed multiple times (once per year).\n"
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "## 2. Setting Up the Environment\n",
    "\n",
    "First, we need to install and import the necessary libraries. You can install Pandas by running:\n",
    "\n",
    "```bash\n",
    "pip install pandas\n",
    "```\n",
    "\n",
    "Now let's import Pandas.\n"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 1,
   "metadata": {},
   "outputs": [],
   "source": [
    "import pandas as pd"
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "## 3. Creating a Panel DataFrame\n",
    "\n",
    "We can create a sample panel dataset. Let's say we have data on three students over three years.\n"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 2,
   "metadata": {},
   "outputs": [],
   "source": [
    "# Creating a sample panel dataset\n",
    "data = {\n",
    "    'Student': ['Alice', 'Alice', 'Alice', 'Bob', 'Bob', 'Bob', 'Charlie', 'Charlie', 'Charlie'],\n",
    "    'Year': [2021, 2022, 2023, 2021, 2022, 2023, 2021, 2022, 2023],\n",
    "    'Score': [85, 90, 95, 80, 85, 88, 78, 82, 85]\n",
    "}\n",
    "\n",
    "panel_data = pd.DataFrame(data)\n",
    "panel_data"
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "## 4. Exploring the Panel Data\n",
    "\n",
    "Let's take a look at our dataset. We can see the scores for each student across different years.\n"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 3,
   "metadata": {},
   "outputs": [],
   "source": [
    "# Displaying the panel data\n",
    "panel_data"
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "## 5. Analyzing Panel Data\n",
    "\n",
    "We can perform various analyses with panel data, such as calculating the average score for each student over the years.\n"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 4,
   "metadata": {},
   "outputs": [],
   "source": [
    "# Calculating the average score for each student\n",
    "average_scores = panel_data.groupby('Student')['Score'].mean()\n",
    "average_scores"
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "## 6. Visualizing Panel Data\n",
    "\n",
    "Visualization helps us understand data better. Let's plot the scores for each student over the years using Matplotlib.\n",
    "\n",
    "You can install Matplotlib by running:\n",
    "\n",
    "```bash\n",
    "pip install matplotlib\n",
    "```\n"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 5,
   "metadata": {},
   "outputs": [],
   "source": [
    "import matplotlib.pyplot as plt\n",
    "\n",
    "# Plotting the scores\n",
    "for student in panel_data['Student'].unique():\n",
    "    subset = panel_data[panel_data['Student'] == student]\n",
    "    plt.plot(subset['Year'], subset['Score'], marker='o', label=student)\n",
    "\n",
    "plt.title('Students\' Scores Over Years')\n",
    "plt.xlabel('Year')\n",
    "plt.ylabel('Score')\n",
    "plt.legend()\n",
    "plt.grid()\n",
    "plt.show()"
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "## Conclusion\n",
    "\n",
    "Panel data is a powerful way to analyze changes over time across multiple subjects. With Pandas, you can easily manipulate and analyze this type of data. Now that you've learned the basics, you can explore more advanced topics like fixed effects and random effects models in panel data analysis!"
   ]
  }
 ],
 "metadata": {
  "kernelspec": {
   "display_name": "Python 3",
   "language": "python",
   "name": "python3"
  },
  "language_info": {
   "codemirror_mode": {
    "name": "ipython",
    "version": 3
   },
   "file_extension": ".py",
   "mimetype": "text/x-python",
   "name": "python",
   "nbconvert_exporter": "python",
   "pygments_lexer": "ipython3",
   "version": "3.8.5"
  }
 },
 "nbformat": 4,
 "nbformat_minor": 4
}
