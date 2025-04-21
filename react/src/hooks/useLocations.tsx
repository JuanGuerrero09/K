import { useState, useEffect } from "react";
import { Location } from "@/types/location";

const useLocations = () => {
  const [locations, setLocations] = useState<Location[]>([]);

  useEffect(() => {
    fetch("/madrid_data.json")
      .then((response) => {
        console.log("📦 Response:", response);
        if (!response.ok) throw new Error("❌ No se pudo cargar el JSON");
        return response.json();
      })
      .then((data: Location[]) => {
        console.log("📊 Datos cargados:", data);
        setLocations(data);
      })
      .catch((error) => console.error("🔥 Error cargando datos: ", error));
  }, []);

  return locations;
};

export default useLocations;
