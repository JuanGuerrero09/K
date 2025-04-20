import React from "react";
import { Marker, Popup } from "react-leaflet";
import { lockedIcon, visitedIcon, unvisitedIcon } from "../../lib/icons";

interface Props {
  places: Location[];
}

const Markers: React.FC<Props> = ({ places }) => {
  function selectIcon(location) {
    const { is_complete, is_unlocked } = location;
    if (is_complete && is_unlocked) return visitedIcon;
    if (!is_complete && is_unlocked) return unvisitedIcon;
    return lockedIcon;
  }

  return (
    <>
      {places.map((location, index) => {
        const geoCode = location.geoCode;

        if (
          !geoCode ||
          geoCode.length !== 2 ||
          geoCode.some((coord) => coord === undefined || coord === null)
        ) {
          console.warn("Ubicación sin coordenadas válidas:", location);
          return null;
        }

        const [lat, lng] = geoCode;

        return (
          <Marker
            key={index}
            position={{ lat, lng }}
            icon={selectIcon(location)}
          >
            <Popup>
              <h2>{location.site_name}</h2>
            </Popup>
          </Marker>
        );
      })}
    </>
  );
};

export default Markers;
