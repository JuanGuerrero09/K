import { useState, useEffect } from "react";
import { Location } from "@/types/location";

const useLocations = () => {
  const [locations, setLocations] = useState<Location[]>([]);

  useEffect(() => {
    // Asumiendo que el archivo JSON está en la carpeta /data
    fetch("./madrid_data.json")
      .then((response) => {
        if (!response.ok) throw new Error("No se pudo cargar el JSON");
        return response.json();
      })
      .then((data: Location[]) => {
        setLocations(data);
      })
      .catch((error) => console.error("Error cargando datos: ", error));
  }, []);

  return locations;
};

export default useLocations;
