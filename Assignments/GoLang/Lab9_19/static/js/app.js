// DOM elements - References to UI elements that will be manipulated
const gameForm = document.getElementById('game-form');
const fetchSequentialBtn = document.getElementById('fetch-sequential');
const fetchConcurrentBtn = document.getElementById('fetch-concurrent');
const fetchLeakyBuffersBtn = document.getElementById('fetch-leaky-buffers');
const resultsSection = document.getElementById('results-section');
const gameResults = document.getElementById('game-results');
const timingValue = document.getElementById('timing-value');
const timingBar = document.querySelector('.progress');
const fetchMethod = document.getElementById('fetch-method');
const gameCheckboxes = document.getElementById('game-checkboxes');
const jsonFileInput = document.getElementById('jsonFile');
const importStatus = document.getElementById('import-status');

// Goroutine visualization elements
const goroutineVisualization = document.getElementById('goroutine-visualization');
const goroutinesContainer = document.getElementById('goroutines-container');
const channelBuffer = document.getElementById('channel-buffer');
const mainResults = document.getElementById('main-results');

// Leaky buffer visualization elements
const leakyBufferVisualization = document.getElementById('leaky-buffer-visualization');
const bufferSlots = document.getElementById('buffer-slots');
const goroutinesUsingBuffers = document.getElementById('goroutines-using-buffers');
const returnedBuffers = document.getElementById('returned-buffers');
const buffersCreated = document.getElementById('buffers-created');
const buffersReused = document.getElementById('buffers-reused');
const memorySaved = document.getElementById('memory-saved');

// Maximum expected time for progress bar calculation (milliseconds)
// Used to normalize the progress bar display
const MAX_EXPECTED_TIME = 3000;

// Buffer pool configuration - should match the Go backend
const BUFFER_POOL_SIZE = 10;
const BUFFER_SIZE_KB = 4; // 4KB per buffer

// Initialize the app when DOM is loaded
document.addEventListener('DOMContentLoaded', initializeApp);

// Event listeners - Set up button click handlers for different fetching methods
fetchSequentialBtn.addEventListener('click', () => fetchGames('sequential'));
fetchConcurrentBtn.addEventListener('click', () => fetchGames('concurrent'));
fetchLeakyBuffersBtn.addEventListener('click', () => fetchGames('leaky-buffers'));

// JSON file input change handler - automatically import when a file is selected
jsonFileInput.addEventListener('change', handleJSONImport);

// Initialize the application
async function initializeApp() {
    // Fetch all games and populate the checkboxes
    await loadGames();
}

// Load games from the server and populate the game checkboxes
async function loadGames() {
    try {
        const response = await fetch('/api/games');
        if (!response.ok) {
            throw new Error(`Server error: ${response.status}`);
        }

        const games = await response.json();
        
        // Clear existing checkboxes
        gameCheckboxes.innerHTML = '';
        
        // Create checkboxes for each game
        games.forEach(game => {
            const checkboxItem = document.createElement('div');
            checkboxItem.className = 'checkbox-item';
            
            checkboxItem.innerHTML = `
                <input type="checkbox" id="game${game.id}" name="ids" value="${game.id}" checked>
                <label for="game${game.id}">${game.title}</label>
            `;
            
            gameCheckboxes.appendChild(checkboxItem);
        });
        
        return games;
    } catch (error) {
        console.error('Error loading games:', error);
        gameCheckboxes.innerHTML = `<div class="error">Error loading games: ${error.message}</div>`;
        return [];
    }
}

// Handle JSON file import - automatically triggered when a file is selected
async function handleJSONImport(event) {
    // Reset status message
    importStatus.innerHTML = '';
    importStatus.className = 'status-message';
    
    // Get the selected file
    const jsonFile = event.target.files[0];
    if (!jsonFile || jsonFile.size === 0) {
        setImportStatus('Please select a JSON file', 'error');
        return;
    }
    
    // Check file extension
    if (!jsonFile.name.toLowerCase().endsWith('.json')) {
        setImportStatus('Please upload a valid JSON file', 'error');
        event.target.value = ''; // Clear the file input
        return;
    }
    
    try {
        // Set loading status
        setImportStatus('Importing games...', '');
        
        // Create FormData and add the file
        const formData = new FormData();
        formData.append('jsonFile', jsonFile);
        
        // Send the file to the server
        const response = await fetch('/api/games/import', {
            method: 'POST',
            body: formData
        });
        
        if (!response.ok) {
            const errorData = await response.text();
            throw new Error(errorData || `Server error: ${response.status}`);
        }
        
        const result = await response.json();
        
        // Show success message
        setImportStatus(result.message, 'success');
        
        // Reload the games to update the checkboxes
        await loadGames();
        
        // Clear the file input for future imports
        event.target.value = '';
        
    } catch (error) {
        console.error('Error importing JSON:', error);
        setImportStatus(`Error: ${error.message}`, 'error');
        event.target.value = ''; // Clear the file input
    }
}

// Set import status message with appropriate styling
function setImportStatus(message, type) {
    importStatus.textContent = message;
    importStatus.className = 'status-message';
    if (type) {
        importStatus.classList.add(type);
    }
}

// Main function to fetch games - handles both sequential and concurrent requests
// This function will call the appropriate Go backend endpoint based on the method
async function fetchGames(method) {
    // Get selected game IDs from the form
    const formData = new FormData(gameForm);
    const selectedIds = formData.getAll('ids');
    
    // Validate at least one game is selected
    if (selectedIds.length === 0) {
        alert('Please select at least one game to fetch');
        return;
    }
    
    // Show loading state and disable buttons during fetch
    fetchSequentialBtn.disabled = true;
    fetchConcurrentBtn.disabled = true;
    fetchLeakyBuffersBtn.disabled = true;
    gameResults.innerHTML = '<div class="loading">Fetching games...</div>';
    resultsSection.classList.remove('hidden');
    
    // Reset visualizations
    resetVisualizations();
    
    // Show appropriate visualization based on method
    if (method === 'leaky-buffers') {
        // Show leaky buffer visualization
        leakyBufferVisualization.classList.remove('hidden');
        goroutineVisualization.classList.add('hidden');
        
        // Initialize and start the leaky buffer visualization
        initBufferPoolVisualization();
        startLeakyBufferVisualization(selectedIds);
    } else {
        // Show goroutine visualization for sequential and concurrent
        goroutineVisualization.classList.remove('hidden');
        leakyBufferVisualization.classList.add('hidden');
        
        if (method === 'concurrent') {
            startConcurrentVisualization(selectedIds);
        } else {
            startSequentialVisualization(selectedIds);
        }
    }
    
    // Update method display to show which approach is being used
    if (method === 'sequential') {
        fetchMethod.textContent = 'Sequential';
    } else if (method === 'concurrent') {
        fetchMethod.textContent = 'Concurrent';
    } else {
        fetchMethod.textContent = 'Leaky Buffers';
    }
    
    try {
        // Build URL with selected IDs for the API call
        const url = `/api/games/${method}?${selectedIds.map(id => `ids=${id}`).join('&')}`;
        
        // Fetch data from server
        const response = await fetch(url);
        
        if (!response.ok) {
            throw new Error(`Server error: ${response.status}`);
        }
        
        // Parse the JSON response from the server
        const data = await response.json();
        
        // Display results in the UI
        displayResults(data);
        
        // Update buffer usage stats if using leaky buffers
        if (method === 'leaky-buffers' && data.BuffersReused) {
            updateBufferStats(data.BuffersReused);
        }
    } catch (error) {
        console.error('Error fetching games:', error);
        gameResults.innerHTML = `<div class="error">Error: ${error.message}</div>`;
    } finally {
        // Reset buttons when operation completes (success or failure)
        fetchSequentialBtn.disabled = false;
        fetchConcurrentBtn.disabled = false;
        fetchLeakyBuffersBtn.disabled = false;
    }
}

// Reset all visualizations to their initial state
function resetVisualizations() {
    // Reset goroutine visualization
    goroutinesContainer.innerHTML = '';
    channelBuffer.innerHTML = '';
    mainResults.innerHTML = '';
    
    // Reset the goroutine visualization title
    const visualizationTitle = document.querySelector('.goroutine-visualization h2');
    if (visualizationTitle) {
        visualizationTitle.textContent = 'Live Goroutine Execution';
    }
    
    // Reset leaky buffer visualization
    bufferSlots.innerHTML = '';
    goroutinesUsingBuffers.innerHTML = '';
    returnedBuffers.innerHTML = '';
    buffersCreated.textContent = '0';
    buffersReused.textContent = '0';
    memorySaved.textContent = '0 KB';
}

// Start the visualization of sequential execution
function startSequentialVisualization(gameIds) {
    // Update the visualization title
    const visualizationTitle = document.querySelector('.goroutine-visualization h2');
    if (visualizationTitle) {
        visualizationTitle.textContent = 'Live Sequential Execution (One at a Time)';
    }
    
    // Get all games first to have their titles for visualization
    fetch('/api/games')
        .then(response => response.json())
        .then(games => {
            // Create a map of game IDs to titles for easy lookup
            const gameMap = games.reduce((map, game) => {
                map[game.id] = game;
                return map;
            }, {});
            
            // Create goroutines for each selected game
            gameIds.forEach((id, index) => {
                const game = gameMap[id] || { id, title: `Game ${id}` };
                createGoroutine(game, index);
            });
            
            // Simulate the sequential execution with animations
            simulateSequentialExecution(gameIds, gameMap);
        })
        .catch(error => {
            console.error('Error fetching games for visualization:', error);
        });
}

// Start the visualization of concurrent execution
function startConcurrentVisualization(gameIds) {
    // Update the visualization title
    const visualizationTitle = document.querySelector('.goroutine-visualization h2');
    if (visualizationTitle) {
        visualizationTitle.textContent = 'Live Concurrent Execution (All at Once)';
    }
    
    // Get all games first to have their titles for visualization
    fetch('/api/games')
        .then(response => response.json())
        .then(games => {
            // Create a map of game IDs to titles for easy lookup
            const gameMap = games.reduce((map, game) => {
                map[game.id] = game;
                return map;
            }, {});
            
            // Create goroutines for each selected game
            gameIds.forEach((id, index) => {
                const game = gameMap[id] || { id, title: `Game ${id}` };
                createGoroutine(game, index);
            });
            
            // Simulate the concurrent execution with animations
            simulateConcurrentExecution(gameIds, gameMap);
        })
        .catch(error => {
            console.error('Error fetching games for visualization:', error);
        });
}

// Initialize the buffer pool visualization with empty slots
function initBufferPoolVisualization() {
    // Clear any existing buffer slots
    bufferSlots.innerHTML = '';
    
    // Create buffer slots (initially empty)
    for (let i = 0; i < BUFFER_POOL_SIZE; i++) {
        const bufferSlot = document.createElement('div');
        bufferSlot.className = 'buffer-slot empty';
        bufferSlot.id = `buffer-slot-${i}`;
        bufferSlot.textContent = i + 1;
        bufferSlots.appendChild(bufferSlot);
    }
    
    // Reset stats
    buffersCreated.textContent = '0';
    buffersReused.textContent = '0';
    memorySaved.textContent = '0 KB';
}

// Start the leaky buffer visualization
function startLeakyBufferVisualization(gameIds) {
    // Get all games first to have their titles for visualization
    fetch('/api/games')
        .then(response => response.json())
        .then(games => {
            // Create a map of game IDs to titles for easy lookup
            const gameMap = games.reduce((map, game) => {
                map[game.id] = game;
                return map;
            }, {});
            
            // Simulate the leaky buffer pattern with animations
            simulateLeakyBuffers(gameIds, gameMap);
        })
        .catch(error => {
            console.error('Error fetching games for visualization:', error);
        });
}

// Simulate the leaky buffer pattern
function simulateLeakyBuffers(gameIds, gameMap) {
    // Keeps track of which buffers are currently in use
    const bufferUsage = new Array(BUFFER_POOL_SIZE).fill(false);
    
    // Keeps track of created and reused buffers
    let created = 0;
    let reused = 0;
    
    // Process each game with a slight delay between them
    gameIds.forEach((id, index) => {
        const game = gameMap[id] || { id, title: `Game ${id}` };
        
        // Simulate getting a buffer from the pool
        setTimeout(() => {
            // Try to find an available buffer
            let bufferIndex = bufferUsage.findIndex(used => !used);
            let isReused = bufferIndex !== -1;
            
            // If no buffer is available, create a new one
            if (bufferIndex === -1) {
                // All buffers are in use, so we'll create a new one
                // (which would be "leaked" when returned)
                bufferIndex = created % BUFFER_POOL_SIZE;
                isReused = false;
                created++;
                buffersCreated.textContent = created;
            } else {
                // Reusing an existing buffer
                reused++;
                buffersReused.textContent = reused;
                
                // Calculate memory saved (4KB per buffer reused)
                const saved = reused * BUFFER_SIZE_KB;
                memorySaved.textContent = `${saved} KB`;
            }
            
            // Mark this buffer as being used
            bufferUsage[bufferIndex] = true;
            
            // Update buffer slot appearance
            const bufferSlot = document.getElementById(`buffer-slot-${bufferIndex}`);
            if (bufferSlot) {
                bufferSlot.className = `buffer-slot ${isReused ? 'reused' : 'new'} being-used`;
            }
            
            // Create a goroutine with buffer element
            const goroutineWithBuffer = document.createElement('div');
            goroutineWithBuffer.className = 'goroutine-with-buffer';
            goroutineWithBuffer.id = `goroutine-buffer-${id}`;
            
            goroutineWithBuffer.innerHTML = `
                <div class="buffer-in-use ${isReused ? 'reused' : 'new'}">${bufferIndex + 1}</div>
                <div class="goroutine-label">${game.title}</div>
            `;
            
            goroutinesUsingBuffers.appendChild(goroutineWithBuffer);
            
            // Simulate processing time
            const processingTime = 100 + (id * 50);
            
            // After processing, return the buffer to the pool
            setTimeout(() => {
                // Remove goroutine with buffer
                goroutineWithBuffer.style.opacity = '0';
                setTimeout(() => {
                    goroutineWithBuffer.remove();
                }, 500);
                
                // Create a returning buffer element
                const returnedBuffer = document.createElement('div');
                returnedBuffer.className = `returned-buffer ${isReused ? 'reused' : 'new'}`;
                returnedBuffer.textContent = bufferIndex + 1;
                returnedBuffers.appendChild(returnedBuffer);
                
                // Mark buffer as available again
                bufferUsage[bufferIndex] = false;
                
                // Update buffer slot appearance
                if (bufferSlot) {
                    bufferSlot.className = `buffer-slot ${isReused ? 'reused' : 'new'}`;
                    setTimeout(() => {
                        // After a delay, clean up the returned buffer element
                        returnedBuffer.remove();
                    }, 1000);
                }
            }, processingTime);
        }, index * 200); // Stagger the start time of each game's processing
    });
}

// Simulate sequential execution with animations - one at a time
function simulateSequentialExecution(gameIds, gameMap) {
    // Process one game at a time, waiting for each to complete before starting the next
    processNextGame(gameIds, gameMap, 0, 0);
}

// Process games one by one in sequence
function processNextGame(gameIds, gameMap, index, totalTime) {
    // Stop if we've processed all games
    if (index >= gameIds.length) return;
    
    const id = gameIds[index];
    const game = gameMap[id] || { id, title: `Game ${id}` };
    
    const goroutine = document.getElementById(`goroutine-${id}`);
    if (!goroutine) return;
    
    // Start processing animation for the current goroutine
    goroutine.classList.add('processing');
    
    // Calculate a simulated processing time (different for each goroutine)
    const processingTime = 100 + (id * 50);
    
    // After the processing time, animate sending the result to the channel
    setTimeout(() => {
        // Mark goroutine as completed
        goroutine.classList.remove('processing');
        goroutine.classList.add('completed');
        
        // After a short delay, show the item in the channel
        setTimeout(() => {
            const colorIndex = (index % 8) + 1;
            
            // Create buffer item in channel
            const bufferItem = document.createElement('div');
            bufferItem.className = `buffer-item goroutine-${colorIndex}`;
            bufferItem.textContent = id;
            channelBuffer.appendChild(bufferItem);
            
            // After the buffer item appears, simulate it being received by the main goroutine
            setTimeout(() => {
                // Create a result item in the main goroutine
                const resultItem = document.createElement('div');
                resultItem.className = 'result-item';
                resultItem.textContent = game.title;
                
                // Add the result to the main goroutine with a slight delay
                setTimeout(() => {
                    mainResults.appendChild(resultItem);
                    
                    // Remove the buffer item after it's been "consumed"
                    bufferItem.style.opacity = '0';
                    setTimeout(() => {
                        bufferItem.remove();
                        
                        // Process the next game in sequence
                        processNextGame(gameIds, gameMap, index + 1, totalTime + processingTime);
                    }, 300);
                }, 100);
            }, 500);
        }, 200);
    }, processingTime);
}

// Create a visual goroutine element
function createGoroutine(game, index) {
    const goroutine = document.createElement('div');
    
    // Use modulo to cycle through available colors if there are more than 8 games
    const colorIndex = (index % 8) + 1;
    
    goroutine.className = `goroutine goroutine-${colorIndex}`;
    goroutine.id = `goroutine-${game.id}`;
    
    goroutine.innerHTML = `
        <div class="goroutine-id">Goroutine ${index + 1}</div>
        <div class="goroutine-title">${game.title}</div>
    `;
    
    // Add a small delay before showing each goroutine
    setTimeout(() => {
        goroutinesContainer.appendChild(goroutine);
    }, index * 100);
}

// Simulate concurrent execution with animations
function simulateConcurrentExecution(gameIds, gameMap) {
    // For each game ID, simulate its processing and completion
    gameIds.forEach((id, index) => {
        const game = gameMap[id] || { id, title: `Game ${id}` };
        
        // Wait for the goroutine to be created
        setTimeout(() => {
            const goroutine = document.getElementById(`goroutine-${id}`);
            if (!goroutine) return;
            
            // Start processing animation
            goroutine.classList.add('processing');
            
            // Calculate a simulated processing time (different for each goroutine)
            // This matches the logic in the Go backend where each game has a different delay
            const processingTime = 100 + (id * 50);
            
            // After the processing time, animate sending the result to the channel
            setTimeout(() => {
                // Mark goroutine as completed and animate it
                goroutine.classList.remove('processing');
                goroutine.classList.add('completed');
                
                // Create a buffer item in the channel after a short delay
                setTimeout(() => {
                    const bufferItem = document.createElement('div');
                    
                    // Use the same color as the goroutine
                    const colorIndex = (index % 8) + 1;
                    
                    bufferItem.className = `buffer-item goroutine-${colorIndex}`;
                    bufferItem.textContent = id;
                    channelBuffer.appendChild(bufferItem);
                    
                    // After the buffer item appears, simulate it being received by the main goroutine
                    setTimeout(() => {
                        // Create a result item in the main goroutine
                        const resultItem = document.createElement('div');
                        resultItem.className = 'result-item';
                        resultItem.textContent = game.title;
                        
                        // Add the result to the main goroutine with a slight delay
                        setTimeout(() => {
                            mainResults.appendChild(resultItem);
                            
                            // Remove the buffer item after it's been "consumed" by the main goroutine
                            bufferItem.style.opacity = '0';
                            setTimeout(() => {
                                bufferItem.remove();
                            }, 300);
                        }, 100);
                    }, 500);
                }, 200);
            }, processingTime);
        }, index * 100 + 300); // Delay to ensure goroutines are created first
    });
}

// Update buffer statistics based on server response
function updateBufferStats(reused) {
    const created = BUFFER_POOL_SIZE - reused > 0 ? BUFFER_POOL_SIZE - reused : 0;
    buffersCreated.textContent = created;
    buffersReused.textContent = reused;
    
    // Calculate memory saved (4KB per buffer reused)
    const saved = reused * BUFFER_SIZE_KB;
    memorySaved.textContent = `${saved} KB`;
}

// Display results in the UI - shows the timing difference between methods
function displayResults(data) {
    // Clear previous results
    gameResults.innerHTML = '';
    
    // Update timing display
    const timeInMs = formatDuration(data.Timing);
    timingValue.textContent = timeInMs;
    
    // Update progress bar
    const progressPercentage = Math.min((parseFloat(data.Timing) / MAX_EXPECTED_TIME) * 100, 100);
    timingBar.style.width = `${progressPercentage}%`;
    
    // Sort games by ID for consistent display
    const sortedGames = [...data.Games].sort((a, b) => a.id - b.id);
    
    // Create game cards for each fetched game
    sortedGames.forEach(game => {
        const card = createGameCard(game);
        gameResults.appendChild(card);
    });
    
    // If no games were returned
    if (sortedGames.length === 0) {
        gameResults.innerHTML = '<div class="no-results">No games found</div>';
    }
}

// Create HTML for a game card - displays each game's information
function createGameCard(game) {
    const card = document.createElement('div');
    card.className = 'game-card';
    
    // Format fetch time if available
    let fetchTimeHtml = '';
    if (game.fetchTime) {
        fetchTimeHtml = `<span class="fetch-time">${formatDuration(game.fetchTime)}</span>`;
    }
    
    // Format buffer badge if applicable (for leaky buffer demo)
    let bufferBadgeHtml = '';
    if (game.bufferReused !== undefined) {
        const badgeClass = game.bufferReused ? 'reused' : 'new';
        const badgeText = game.bufferReused ? 'Reused Buffer' : 'New Buffer';
        bufferBadgeHtml = `<span class="buffer-badge ${badgeClass}">${badgeText}</span>`;
    }
    
    // Format multiplayer badge if applicable
    const multiplayerBadge = game.isMultiplayer 
        ? '<span class="multiplayer-badge">Multiplayer</span>' 
        : '';
    
    // Create the card content
    card.innerHTML = `
        ${fetchTimeHtml}
        ${bufferBadgeHtml}
        <h3 class="title">${game.title}</h3>
        <div class="game-detail">
            <span class="detail-label">Genre:</span>
            <span>${game.genre}</span>
        </div>
        <div class="game-detail">
            <span class="detail-label">Publisher:</span>
            <span>${game.publisher}</span>
        </div>
        <div class="game-detail">
            <span class="detail-label">Year:</span>
            <span>${game.releaseYear}</span>
        </div>
        <div class="game-detail">
            <span class="detail-label">Rating:</span>
            <span>${game.rating}/10</span>
        </div>
        <div class="game-detail">
            <span class="detail-label">Price:</span>
            <span>$${game.price.toFixed(2)}</span>
        </div>
        ${multiplayerBadge}
    `;
    
    return card;
}

// Format duration from nanoseconds to a readable string
// Go returns durations in nanoseconds, so we need to convert
function formatDuration(nanoseconds) {
    // Convert nanoseconds to milliseconds
    const ms = nanoseconds / 1000000;
    
    if (ms < 1000) {
        return `${ms.toFixed(2)}ms`;
    } else {
        return `${(ms / 1000).toFixed(2)}s`;
    }
} 