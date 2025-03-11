import { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar/Sidebar";
import Map from "../components/Map/Map";
import useLocations from "../hooks/useLocations";
import Navbar from "../components/Navbar/Navbar";
import Hero from "../components/Hero/Hero";

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
    <>
      <div className="flex flex-col items-center min-h-screen bg-gray-100 drawer ">
        <input id="my-drawer-4" type="checkbox" className="drawer-toggle" />
        {/* Contenedor del texto */}
        <Navbar></Navbar>
        <div className="bg-white shadow-md rounded-lg p-6 w-full text-center drawer-content">
          <Hero></Hero>
          <p className="mt-4 text-lg font-medium text-gray-700">
            Tu puntuación actual:{points}
            <span className="text-blue-600">{points}</span>
          </p>
          {/* Mapa */}
          <div className="mt-6 w-full flex justify-center">
            <Map />
          </div>
        </div>
        <div className="drawer-side">
          <Sidebar></Sidebar>
        </div>
      </div>
    </>
  );
}
