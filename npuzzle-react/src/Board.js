import React, { useState } from "react";
import Tile from "./Tile";
import {PUZZLES_COUNT, BOARD_SIZE, REQ_WAIT_CALC} from "./constants"
import { canSwap, shuffle, swap, isSolved } from "./helpers"
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import Box, { BoxProps } from '@mui/material/Box';
import Grid from "@mui/material/Grid";
import Tooltip from "@mui/material/Tooltip"
import Requests from "./api/Requests";
import ArrowForwardIcon from '@mui/icons-material/ArrowForward';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import RestartAltIcon from '@mui/icons-material/RestartAlt';

import FormGroup from '@mui/material/FormGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import Switch from '@mui/material/Switch';
import useForceUpdate from 'use-force-update';


function Item(props) {
  const { sx, ...other } = props;
  return (
      <Box
          sx={{
            p: 1,
            m: 1,
            ...sx,
          }}
          {...other}
      />
  );
}

function Board() {

  const [tiles, setTiles] = useState([...Array(PUZZLES_COUNT[0]).keys()]);
  const [isStarted, setIsStarted] = useState(false);
  const [isStopCalc, setStopCalc] = useState(false);
  const [isTileClickRun, setTileClickRun] = useState(false);
  const [isCurrentEmptyTileIndex, setCurrentEmptyTileIndex] = useState(0);
  const [isCurrentVerbs, setCurrentVerbs] = useState("");
  const [isCurrentVerbIndex, setCurrentVerbIndex] = useState(0);
  const [isManhattanHe, setManhattanHe] = useState(true);
  const [isEuclideanHe, setEuclideanHe] = useState(false);
  const [isHammingHe, setHammingHe] = useState(false);
  const [isSolvable, setSolvable] = useState(true);



  // const forceUpdate = useForceUpdate();

  let shuffledTiles;
  let countGetWaitReq;

  const setTileClickRegime = (value) => {
    setTileClickRun(value)
  }

  const setWaitTimeout = (num) => {
    REQ_WAIT_CALC[0] = num
    setTiles([...Array(PUZZLES_COUNT[0]).keys()])
    setIsStarted(false)
    setStopCalc(false)
  }

  const setTilesToNum = (num) => {
    PUZZLES_COUNT[1] = num
    PUZZLES_COUNT[0] = num * num
    setTiles([...Array(PUZZLES_COUNT[0]).keys()])
    setIsStarted(false)
    setStopCalc(false)
  }

  const handleGetRequest = () => {
    if (countGetWaitReq === REQ_WAIT_CALC[0] || countGetWaitReq > REQ_WAIT_CALC[0]) {
      return; } else {
      Requests.getPath().then(res => {
        if (res.status === 200) {
          console.log('get success path');
          console.log(res.data)
          setTiles(shuffledTiles)
          handleSuccessPath(shuffledTiles, PUZZLES_COUNT[0] - 3, res.data)
          return;
        } else {
          console.log('wait for get path');
          setTimeout(() => {
            handleGetRequest();
          }, 1000);
          countGetWaitReq++;
          if (countGetWaitReq === REQ_WAIT_CALC[0]) {
            Requests.getCalculating()
            setStopCalc(true)
            console.log('stop get path');
            return;
          }
          return;
        }
      })
    }
  }

  const GetServerTiles = async () => {
    let he;

    if (isManhattanHe) {
      he = "Manhattan";
    } else if (isEuclideanHe) {
      he = "Euclidean";
    } else if (isHammingHe) {
      he = "Hamming";
    }
    if (isSolvable) {
      console.log(isSolvable);
    }

    console.log(he)
    const res = await Requests.getState(PUZZLES_COUNT[0], he, isSolvable);
    if  (res.status === 200) {
      console.log(res.data)
      shuffledTiles = res.data;
      setTiles(shuffledTiles);
      setStopCalc(false);
      countGetWaitReq = 0;
    //   setTimeout(() => {
    //     handleGetRequest(shuffledTiles);
    //   }, 1000);
    // } else {
    //   console.log(res.status);
    }
  }

  // const shuffleTiles = async () => {
  //   shuffledTiles = shuffle(tiles)
  //   setTiles(shuffledTiles);
  //   console.log(shuffledTiles)
  //   setStopCalc(false);
  //
  //   const res = await Requests.putState({
  //     inputState: shuffledTiles,
  //     emptyTile: [PUZZLES_COUNT[0] - 3]
  //   })
  //   if (res.status === 200) {
  //     Requests.startAlgo()
  //     countGetWaitReq = 0;
  //     setTimeout(() => {
  //       handleGetRequest(shuffledTiles);
  //     }, 1000);
  //   }
  // }

  const swapTiles = (tileIndex) => {
    if (canSwap(tileIndex, tiles.indexOf(tiles.length - 3))) {
      const swappedTiles = swap(tiles, tileIndex, tiles.indexOf(tiles.length - 3))
      setTiles(swappedTiles)
      // console.log("success")
    } else {
      // console.log("smth wrong")
    }
  }

  const handleTileClick = (index, tile) => {
    if (isTileClickRun) {
      swapTiles(index)
    }
  }

  const swapTilesAutoClick = (startTileIndex, endTileIndex, tiles) => {
    let swappedTiles;
    if (canSwap(startTileIndex, endTileIndex)) {
      swappedTiles = swap(tiles, startTileIndex, endTileIndex)
      setTiles(swappedTiles)
    }
  }

  const swapTilesAuto = (startTileIndex, endTileIndex, shuffledTiles) => {
    let swappedTiles;
    if (canSwap(startTileIndex, endTileIndex)) {
      swappedTiles = swap(shuffledTiles, startTileIndex, endTileIndex)
      setTiles(swappedTiles)
      // console.log("success")
    } else {
      // console.log("smth wrong")
    }
    return swappedTiles
  }

  function sleep(milliseconds) {
    const date = Date.now();
    let currentDate = null;
    do {
      currentDate = Date.now();
    } while (currentDate - date < milliseconds);
  }

  const handleSuccessPath = (shuffledTiles, emptyTile, verbs) => {
    let startEmptyTileIndex;

    for (let i = 0; i < shuffledTiles.length; i++) {
      if (shuffledTiles[i] === emptyTile) {
        startEmptyTileIndex = i;
      }
    }
    setCurrentEmptyTileIndex(startEmptyTileIndex);
    setCurrentVerbs(verbs["successPath"]);
    setCurrentVerbIndex(0);

  }

  const SwitchTilesToEnd = () => {
    let index = isCurrentEmptyTileIndex;
    let verbIndex = isCurrentVerbIndex;
    if (verbIndex > -1 && verbIndex < isCurrentVerbs.length - 1) {
      let shuffledTiles = tiles;
      let i;
      for (i = verbIndex; i < isCurrentVerbs.length - 1; i++) {
        if (isCurrentVerbs[i] === "LEFT") {
          index--;
          shuffledTiles = swapTilesAuto(index, shuffledTiles.indexOf(shuffledTiles.length - 3), shuffledTiles)
        } else if (isCurrentVerbs[i] === "RIGHT") {
          index++;
          shuffledTiles = swapTilesAuto(index, shuffledTiles.indexOf(shuffledTiles.length - 3), shuffledTiles)
        } else if (isCurrentVerbs[i] === "UP") {
          index = index - PUZZLES_COUNT[1];
          shuffledTiles = swapTilesAuto(index, shuffledTiles.indexOf(shuffledTiles.length - 3), shuffledTiles)
        } else if (isCurrentVerbs[i] === "DOWN") {
          index = index + PUZZLES_COUNT[1];
          shuffledTiles = swapTilesAuto(index, shuffledTiles.indexOf(shuffledTiles.length - 3), shuffledTiles)
        }
      }
      setCurrentEmptyTileIndex(index);
      console.log(i)
      if (i > isCurrentVerbs.length ) {
        setCurrentVerbIndex(isCurrentVerbs.length - 1);
      } else {
        setCurrentVerbIndex(i - 1);
      }
    }
  }

  const handleTileSwitchStep = (step) => {
    let index;
    let verbIndex = isCurrentVerbIndex;
    if (verbIndex > -1 && verbIndex < isCurrentVerbs.length) {
      if (step === "back") {
        if (verbIndex === 0) {
          return;
        }
        --verbIndex;
        if (isCurrentVerbs[verbIndex] === "LEFT") {
          index = isCurrentEmptyTileIndex + 1;
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        } else if (isCurrentVerbs[verbIndex] === "RIGHT") {
          index = isCurrentEmptyTileIndex - 1;
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        } else if (isCurrentVerbs[verbIndex] === "UP") {
          index = isCurrentEmptyTileIndex + PUZZLES_COUNT[1];
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        } else if (isCurrentVerbs[verbIndex] === "DOWN") {
          index = isCurrentEmptyTileIndex - PUZZLES_COUNT[1];
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        }
      } else if (step === "forward") {
        if (verbIndex === isCurrentVerbs.length - 1) {
          return;
        }
        if (isCurrentVerbs[verbIndex] === "LEFT") {
          index = isCurrentEmptyTileIndex - 1;
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        } else if (isCurrentVerbs[verbIndex] === "RIGHT") {
          index = isCurrentEmptyTileIndex + 1;
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        } else if (isCurrentVerbs[verbIndex] === "UP") {
          index = isCurrentEmptyTileIndex - PUZZLES_COUNT[1];
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        } else if (isCurrentVerbs[verbIndex] === "DOWN") {
          index = isCurrentEmptyTileIndex + PUZZLES_COUNT[1];
          swapTilesAutoClick(index, tiles.indexOf(tiles.length - 3), tiles);
          setCurrentEmptyTileIndex(index);
        }
        ++verbIndex;
      }
      setCurrentVerbIndex(verbIndex);
    }
    console.log(isCurrentVerbIndex)
  }

  const handleShuffleClick = () => {
    // shuffleTiles()
    GetServerTiles()
  }

  const handleStartClick = () => {
    // shuffleTiles()
    GetServerTiles()
    setIsStarted(true)
  }



  let tileColor;
  if (isCurrentVerbIndex === isCurrentVerbs.length - 1) {
    tileColor = `chartreuse`;
  } else {
    tileColor = `#924fb9`;
  }

  const pieceWidth = Math.round(BOARD_SIZE / PUZZLES_COUNT[1]);
  const pieceHeight = Math.round(BOARD_SIZE / PUZZLES_COUNT[1]);
  const style = {
    width: BOARD_SIZE,
    height: BOARD_SIZE,
  };

  return (
    <>
      <div style={{ width: '100%' }}>
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'space-between',
            p: 1,
            m: 1,
            borderRadius: 1,
          }}
        >
          <Item>
            <Grid
                container
                direction="column"
                justifyContent="space-evenly"
                alignItems="flex-start"
            >
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left-start">
                {PUZZLES_COUNT[0] === 9 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(3)}>3x3 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(3)}>3x3 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left">
                {PUZZLES_COUNT[0] === 16 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(4)}>4x4 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(4)}>4x4 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left-end">
                {PUZZLES_COUNT[0] === 25 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(5)}>5x5 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(5)}>5x5 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left-end">
                {PUZZLES_COUNT[0] === 36 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(6)}>6x6 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(6)}>6x6 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left-end">
                {PUZZLES_COUNT[0] === 49 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(7)}>7x7 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(7)}>7x7 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left-end">
                {PUZZLES_COUNT[0] === 64 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(8)}>8x8 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(8)}>8x8 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left-end">
                {PUZZLES_COUNT[0] === 81 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(9)}>9x9 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(9)}>9x9 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change puzzles count" placement="left-end">
                {PUZZLES_COUNT[0] === 100 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTilesToNum(10)}>10x10 puzzles</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTilesToNum(10)}>10x10 puzzles</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <Tooltip title="change solvability" placement="left-end">
                {isSolvable ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setSolvable(false)}>solvable</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setSolvable(true)}>solvable</Button>)
                }
              </Tooltip>

            </Grid>
          </Item>
          <Item>
            <ul style={style} className="board">
              {tiles.map((tile, index) => (
                <Tile
                    color={tileColor}
                    key={tile}
                    index={index}
                    tile={tile}
                    width={pieceWidth}
                    height={pieceHeight}
                    handleTileClick={handleTileClick}
                />
              ))}
            </ul>
          </Item>
          <Item>
            <Grid>
              <br />
              <Tooltip title="change wait seconds calculating timeout" placement="right-end">
                {REQ_WAIT_CALC[0] === 3 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setWaitTimeout(3)}>3 sec timeout</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setWaitTimeout(3)}>3 sec timeout</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <br />
              <Tooltip title="change wait seconds calculating timeout" placement="right-end">
                {REQ_WAIT_CALC[0] === 5 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setWaitTimeout(5)}>5 sec timeout</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setWaitTimeout(5)}>5 sec timeout</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <br />
              <Tooltip title="change wait seconds calculating timeout" placement="right-end">
                {REQ_WAIT_CALC[0] === 10 ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setWaitTimeout(10)}>10 sec timeout</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setWaitTimeout(10)}>10 sec timeout</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <br />
              <br />
              <br />
              <Tooltip title="automatically switch tiles" placement="right-end">
                {!isTileClickRun ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTileClickRegime(true)}>server switch</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTileClickRegime(false)}>server switch</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <br />
              <Tooltip title="manually switch tiles" placement="right-end">
                {isTileClickRun ?
                    (<Button style={{color: `chartreuse`}} onClick={() => setTileClickRegime(false)}>manual switch</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => setTileClickRegime(true)}>manual switch</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <br />
              <br />
              <br />
              <Tooltip title="Manhattan heuristics" placement="right-end">
                {isManhattanHe ?
                    (<Button style={{color: `chartreuse`}} onClick={() => {setManhattanHe(true); setEuclideanHe(false); setHammingHe(false)}}>Manhattan he</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => {setManhattanHe(true); setEuclideanHe(false); setHammingHe(false)}}>Manhattan he</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <br />
              <Tooltip title="Euclidean heuristics" placement="right-end">
                {isEuclideanHe ?
                    (<Button style={{color: `chartreuse`}} onClick={() => {setManhattanHe(false); setEuclideanHe(true); setHammingHe(false)}}>Euclidean he</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => {setManhattanHe(false); setEuclideanHe(true); setHammingHe(false)}}>Euclidean he</Button>)
                }
              </Tooltip>
              <br />
              <br />
              <br />
              <Tooltip title="Hamming heuristics" placement="right-end">
                {isHammingHe ?
                    (<Button style={{color: `chartreuse`}} onClick={() => {setManhattanHe(false); setEuclideanHe(false); setHammingHe(true)}}>Hamming he</Button>) :
                    (<Button style={{color: `#924fb9`}} onClick={() => {setManhattanHe(false); setEuclideanHe(false); setHammingHe(true)}}>Hamming he</Button>)
                }
              </Tooltip>
            </Grid>

          </Item>
        </Box>
      </div>
      <Stack direction="row" spacing={1}>
        {!isTileClickRun ?
            (<ArrowBackIcon fontSize="large" style={{color: `#924fb9`}} onClick={() => handleTileSwitchStep("back")}></ArrowBackIcon>) :
            (<ArrowBackIcon fontSize="large" style={{color: "grey"}}></ArrowBackIcon>)
        }
        {!isTileClickRun ?
            (<RestartAltIcon fontSize="large" style={{color: `#924fb9`}} onClick={() => SwitchTilesToEnd()}></RestartAltIcon>) :
            (<RestartAltIcon fontSize="large" style={{color: "grey"}}></RestartAltIcon>)
        }
        {!isTileClickRun ?
            (<ArrowForwardIcon fontSize="large" style={{color: `#924fb9`}} onClick={() => handleTileSwitchStep("forward")}></ArrowForwardIcon>) :
            (<ArrowForwardIcon fontSize="large" style={{color: "grey"}}></ArrowForwardIcon>)
        }
      </Stack>

      <Stack direction="row" spacing={1}>
        {!isStarted ?
            (<Button style={{color: `chartreuse`}} onClick={() => handleStartClick()}>Start</Button>) :
            !isStopCalc ?
                (<Button style={{color: `#924fb9`}} onClick={() => handleShuffleClick()}>Restart</Button>) :
                (<Button style={{color: `#924fb9`}} onClick={() => handleShuffleClick()}>Timeout of calculation. Restart</Button>)
        }
        {/*<FormGroup>*/}
        {/*  <FormControlLabel style={{color: `#924fb9`}} control={<Switch defaultChecked color="secondary"/>} label="automatically switch tiles" />*/}
        {/*</FormGroup>*/}
      </Stack>
    </>
  );
}

export default Board;
