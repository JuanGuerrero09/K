import { useState, useEffect } from "react";

const useLocations = () => {
  const [locations, setLocations] = useState([]);

  useEffect(() => {
    // Asumiendo que el archivo JSON está en la carpeta /data
    fetch("/data/planes_madrid.json")
      .then((response) => response.json())
      .then((data) => setLocations(data));
  }, []);
  console.log(locations);

  return locations;
};

export default useLocations;
