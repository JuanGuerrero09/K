import { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar/Sidebar";
import Map from "../components/Map/Map";
import useLocations from "../hooks/useLocations";
import Navbar from "../components/Navbar/Navbar";

export default function Main() {
  const locations = useLocations();
  const [points, setPoints] = useState(0);

  function calculatePoints(locations) {
    let sum = 0;
    for (const place of locations) {
      if (place.is_complete) {
        sum += place.points;
      }
    }
    setPoints(sum);
  }

  useEffect(() => {
    calculatePoints(locations);
  }, []);

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-100 p-6 drawer drawer-end">
      <Navbar></Navbar>
      <input id="my-drawer-4" type="checkbox" className="drawer-toggle" />
      {/* Contenedor del texto */}
      <div className="bg-white shadow-md rounded-lg p-6 w-full text-center drawer-content">
        <label htmlFor="my-drawer-4" className="drawer-button btn btn-primary">
          Open drawer
        </label>
        <h1 className="text-3xl font-bold text-gray-800 border-b-2 border-gray-300 pb-2">
          Tu guía en <span className="text-blue-600 font-bold">Madrid</span>
        </h1>
        <p className="text-gray-600 mt-1">
          Guía interactiva para descubrir Madrid
        </p>
        <p className="mt-4 text-lg font-medium text-gray-700">
          Tu puntuación actual: <span className="text-blue-600">{points}</span>
        </p>
        {/* Mapa */}
        <div className="mt-6 w-full  flex justify-center">
          <Map className="max-w-3xl" />
        </div>
      </div>
      <div className="drawer-side">
        <Sidebar></Sidebar>
      </div>
    </div>
  );
}
