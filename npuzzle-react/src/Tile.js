import React from "react";
import { Motion, spring } from "react-motion";
import { getMatrixPosition, getVisualPosition } from "./helpers";
import { PUZZLES_COUNT, BOARD_SIZE } from "./constants"

function Tile(props) {
  const { color, tile, index, width, height, handleTileClick } = props;
  const { row, col } = getMatrixPosition(index);
  // console.log("row = ", row);
  // console.log("col = ", col);
  const visualPos = getVisualPosition(row, col, width, height);
  const TILE_COUNT = PUZZLES_COUNT[0];
  const GRID_SIZE = PUZZLES_COUNT[1];
  const tileStyle = {
    // paddingLeft: `${50}px`,
    // paddingRight: `${50}px`,
    // paddingHeight: `${50}px`,
    // paddingBottom: `${50}px`,
    // border: 1px solid #FFD1AA;
    // border: `${1}px solid #FFD1AA`,
    // borderRadius: `${3}px`,
    background: color,
    width: `calc(95% / ${ GRID_SIZE })`,
    height: `calc(95% / ${ GRID_SIZE })`,
    translateX: visualPos.x,
    translateY: visualPos.y,
    backgroundSize: `${ BOARD_SIZE * 1.25}px`,
    backgroundPosition: `${(100 / GRID_SIZE) * (tile % GRID_SIZE)}% ${(100 / GRID_SIZE) * (Math.floor(tile / GRID_SIZE))}%`,

  };
  const motionStyle = {
    translateX: spring(visualPos.x),
    translateY: spring(visualPos.y)
  }

  return (
    <Motion style={motionStyle}>
      {({ translateX, translateY }) => (
        <li
          style={{
            ...tileStyle,
            transform: `translate3d(${translateX}px, ${translateY}px, 0)`,
            // Is last tile?
            opacity: tile === TILE_COUNT - 3 && color !== `chartreuse` ? 0 : 1,
          }}
          className="tile"
          onClick= {() => handleTileClick(index, tile)}
        >
          {`${tile + 1}`}
        </li>
      )}
    </Motion>
  );
}

export default Tile;
