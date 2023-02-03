package function

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"chen.tihe/metadata/contracts"
	"chen.tihe/metadata/config"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}
var rpc string
var cfg *config.Config

func init() {
	// Init config once
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config file", err)
	}
	cfg = config

	if config.Stage == "prod" {
		rpc = cfg.MainnetRpc
	} else if config.Stage == "dev" {
		rpc = cfg.GoerliRpc
	}

	functions.HTTP("CloneX", searchTokenMetadata)
}

func searchTokenMetadata(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Allow Get Method only
	switch r.Method {
	case http.MethodGet:
		// Path variable
		tokenId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/"))
		if err != nil {
			http.Error(w, "tokenId is not the number", http.StatusInternalServerError)
			return
		}
	
		if !isValidTokenId(tokenId) {
			http.Error(w, "Invalid Token Id", http.StatusInternalServerError)
			return
		}

		tokenURI, err := url.JoinPath(cfg.BaseUrl, strconv.Itoa(tokenId))
		if err != nil {
			http.Error(w, "Error making tokenURI", http.StatusInternalServerError)
			return
		}

		resp, err := client.Get(tokenURI)
		if err != nil {
			http.Error(w, "Error making request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, string(body))
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func isValidTokenId(tokenId int) bool {
	// Connect ethclien
	ethclient, err := ethclient.Dial(rpc)
    if err != nil {
        log.Fatal(err)
    }
    address := common.HexToAddress(cfg.ContractAddress)

	// Get the instance of smart contract
    instance, err := clonex.NewClonex(address, ethclient)
    if err != nil {
        log.Fatal(err)
    }

	// Type of currTokenIdBig: bin.Int
    currTokenIdBig, err := instance.TotalSupply(nil)
    if err != nil {
        log.Fatal(err)
    }

	// Convert bin.Int to int
	currTokenIdS := currTokenIdBig.String()
	currTokenId, err := strconv.Atoi(currTokenIdS)
	if err != nil {
		log.Fatal(err)
	}

	if tokenId > currTokenId {
		return false
	}

	return true
}