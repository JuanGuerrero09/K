// import useLocations from "../../hooks/useLocations";

// const getCategories() => {
//
import { Location } from "@/types/location";
import { X, MapPin, CheckCircle } from "lucide-react";

interface SidebarProps {
  isOpen: boolean;
  locations: Location[];
  onLocationSelect: (location: Location) => void;
  onClose: () => void;
}

const Sidebar = ({
  isOpen,
  locations,
  onLocationSelect,
  onClose,
}: SidebarProps) => {
  return (
    <div
      className={`
        absolute inset-y-0 left-0 z-50 w-80 bg-white shadow-lg transform transition-transform duration-300 ease-in-out
        ${isOpen ? "translate-x-0" : "-translate-x-full"}
      `}
    >
      <div className="p-4 border-b flex justify-between items-center">
        <h2 className="text-lg font-bold">Madrid Locations</h2>
        <button
          onClick={onClose}
          className="p-1.5 rounded-full hover:bg-gray-100 transition-colors"
        >
          <X size={20} />
        </button>
      </div>

      <div className="p-4 h-[calc(100vh-70px)] overflow-y-auto">
        <div className="mb-4">
          <div className="bg-blue-50 p-3 rounded-lg border border-blue-100">
            <h3 className="font-medium text-blue-800 mb-1">Your Progress</h3>
            <div className="w-full bg-gray-200 rounded-full h-2.5">
              <div
                className="bg-blue-600 h-2.5 rounded-full transition-all duration-500"
                style={{
                  // width: `${(locations.filter((l) => l.visited).length / locations.length) * 100}%`,
                  width: 100,
                }}
              ></div>
            </div>
            <p className="text-sm text-blue-700 mt-2">
              {locations.filter((l) => l.visited).length} of {locations.length}{" "}
              locations visited
            </p>
          </div>
        </div>

        <ul className="space-y-2">
          {locations.length === 0 ? (
            <p className="text-gray-500">Cargando ubicaciones...</p>
          ) : (
            locations?.map((location) => (
              <li key={location.id}>
                <button
                  onClick={() => {
                    onLocationSelect(location);
                    onClose();
                  }}
                  className={`
                  w-full text-left p-3 rounded-lg flex items-start gap-3 transition-colors
                  ${location.visited ? "bg-green-50 border border-green-100" : "bg-white border hover:bg-gray-50"}
                `}
                >
                  <div
                    className={`
                    p-2 rounded-full mt-0.5
                    ${location.visited ? "bg-green-500 text-white" : "bg-red-500 text-white"}
                  `}
                  >
                    <MapPin size={16} />
                  </div>
                  <div>
                    <div className="flex items-center gap-2">
                      <h3 className="font-medium">{location.name}</h3>
                      {location.visited && (
                        <CheckCircle size={16} className="text-green-500" />
                      )}
                    </div>
                    <p className="text-sm text-gray-600 mt-1">
                      {location.description}
                    </p>
                    <div className="mt-2 text-sm font-medium text-yellow-700">
                      {location.points} points
                    </div>
                  </div>
                </button>
              </li>
            ))
          )}
        </ul>
      </div>
    </div>
  );
};

export default Sidebar;

//
// export default function Sidebar() {
//   return (
//     <>
//       <label
//         htmlFor="my-drawer-4"
//         aria-label="close sidebar"
//         className="drawer-overlay"
//       ></label>
//       <ul className="menu flex bg-base-100 text-base-content min-h-full w-80 p-4">
//         {/* Sidebar content here */}
//         {/* For TSX uncomment the commented types below */}
//         <h2 className="text-3xl mt-4 self-center font-bold text-blue-600">
//           Progress
//         </h2>
//         <div
//           className="radial-progress self-center mt-6"
//           style={{ "--value": 70 } /* as React.CSSProperties */}
//           aria-valuenow={70}
//           role="progressbar"
//         >
//           70%
//         </div>
//         <li>
//           <a>Sidebar Item 1</a>
//         </li>
//         <progress
//           className="progress progress-primary w-56"
//           value={50}
//           max="100"
//         ></progress>
//         <li>
//           <a>Sidebar Item 2</a>
//         </li>
//         <progress
//           className="progress progress-primary w-56"
//           value="70"
//           max="100"
//         ></progress>
//       </ul>
//     </>
//   );
// }
