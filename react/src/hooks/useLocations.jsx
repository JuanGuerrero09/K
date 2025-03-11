import { useState, useEffect } from "react";

/**
 * @typedef {Object} Location
 * @property {string} site_name - Nombre del lugar
 * @property {string} site - Ciudad o ubicación general
 * @property {[number, number]} geoCode - Coordenadas [lat, lng]
 * @property {number} points - Puntos asignados
 * @property {string} category - Categoría del lugar
 * @property {boolean} is_complete - Si está completado
 * @property {string} date_of_completion - Fecha de finalización (si aplica)
 * @property {boolean} is_unlocked - Si está desbloqueado
 */

/**
 * Hook para obtener la lista de ubicaciones
 * @returns {Location[]} Lista de ubicaciones
 */
const useLocations = () => {
  const [locations, setLocations] = useState([]);

  useEffect(() => {
    // Asumiendo que el archivo JSON está en la carpeta /data
    fetch("/data/planes_madrid.json")
      .then((response) => response.json())
      .then((data) => setLocations(data));
  }, []);

  return locations;
};

// {
//   "site_name": "Gran Vía",
//   "site": "Madrid",
//   "geoCode": [40.420177, -3.703928],
//   "points": 1,
//   "category": "planes_madrid",
//   "is_complete": false,
//   "date_of_completion": "",
//   "is_unlocked": true
//

export default useLocations;
