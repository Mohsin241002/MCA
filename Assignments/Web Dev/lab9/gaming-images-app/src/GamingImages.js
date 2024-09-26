import React, { useEffect, useState } from 'react';
import axios from 'axios';

const GamingImages = () => {
  const [images, setImages] = useState([]);
  const [error, setError] = useState(null);

  // Replace with your GitHub repository details
  const owner = 'Mohsin241002';
  const repo = 'ImagesAPI';
  const folderPath = 'img'; // Folder inside repo containing images

  useEffect(() => {
    const fetchImagesFromGitHub = async () => {
      const url = `https://api.github.com/repos/${owner}/${repo}/contents/${folderPath}`;

      try {
        const response = await axios.get(url);
        const imageFiles = response.data.filter((file) =>
          file.name.match(/\.(jpg|jpeg|png|gif)$/i) // Filters to keep only image files
        );
        setImages(imageFiles);
      } catch (error) {
        setError('Error fetching images from GitHub');
        console.error(error);
      }
    };

    fetchImagesFromGitHub();
  }, [owner, repo, folderPath]);

  return (
    <div className="container mx-auto p-8">
      <h1 className="text-center text-4xl font-bold my-8 text-gray-800">
        Gaming Image Gallery
      </h1>

      {error ? (
        <p className="text-red-500">{error}</p>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 p-4">
          {images.map((image, index) => (
            <div key={index} className="max-w-xs rounded overflow-hidden shadow-lg">
              <img
                src={image.download_url}
                alt={image.name}
                className="w-full h-48 object-cover"
              />
              <div className="px-6 py-4">
                <div className="font-bold text-xl mb-2">{image.name}</div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default GamingImages;
