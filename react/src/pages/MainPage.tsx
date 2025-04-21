import { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar/Sidebar";
import Map from "../components/Map/Map";
import useLocations from "../hooks/useLocations";
import Navbar from "../components/Navbar/Navbar";
import { Location } from "@/types/location";

export default function Main() {
  const locations = useLocations();
  const [points, setPoints] = useState(0);
  const [sidebarOpen, setSidebarOpen] = useState(false);

  function calculatePoints(locations: Location[]) {
    const total = locations.reduce(
      (sum, loc) => (loc.visited ? sum + loc.points : sum),
      0,
    );
    setPoints(total);
  }

  useEffect(() => {
    calculatePoints(locations);
  }, [locations]);
  const [selectedLocation, setSelectedLocation] = useState<Location | null>(
    null,
  );

  const handleLocationSelect = (location: Location) => {
    setSelectedLocation(location);
  };

  // const handleLocationVisit = (locationId: number) => {
  //   setLocations((prev) =>
  //     prev.map((loc) => {
  //       if (loc.id === locationId && !loc.visited) {
  //         setTotalPoints((current) => current + loc.points);
  //         return { ...loc, visited: true };
  //       }
  //       return loc;
  //     }),
  //   );
  // };

  return (
    <>
      {/* <div className="flex flex-col items-center min-h-screen bg-gray-100 "> */}
      <div className="flex flex-col h-screen bg-gradient-to-br from-blue-50 to-indigo-50">
        {/* Contenedor del texto */}

        <Navbar
          totalPoints={points}
          onToggleSidebar={() => setSidebarOpen(!sidebarOpen)}
        ></Navbar>

        <div className="flex flex-1 relative overflow-hidden">
          <Sidebar
            isOpen={sidebarOpen}
            locations={locations}
            onLocationSelect={handleLocationSelect}
            onClose={() => setSidebarOpen(false)}
          />

          <main className="flex-1 relative bg-white shadow-md rounded-lg p-6 w-full text-center ">
            {/* <PointsOverlay totalPoints={totalPoints} /> */}
            {/**/}
            {/* {selectedLocation && ( */}
            {/*   <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2 w-full max-w-md px-4"> */}
            {/*     <LocationCard */}
            {/*       location={selectedLocation} */}
            {/*       onVisit={handleLocationVisit} */}
            {/*     /> */}
            {/*   </div> */}
            {/* )} */}
            <h1 className="text-3xl font-bold text-gray-800 border-b-2 border-gray-300 pb-2">
              Tu guía en <span className="text-blue-600 font-bold">Madrid</span>
            </h1>
            <p className="text-gray-600 mt-1">
              Guía interactiva para descubrir Madrid
            </p>
            <div className="bg-white shadow-md rounded-lg p-2 w-full text-center ">
              <p className="mt-4 text-lg font-medium text-gray-700">
                Tu puntuación actual:
                <span className="text-blue-600"> {points}</span>
              </p>
              {/* Mapa */}
              <div className="mt-6 w-full flex justify-center">
                <Map
                  locations={locations}
                  selectedLocation={selectedLocation}
                />
              </div>
            </div>
          </main>
        </div>
      </div>
    </>
  );
}
