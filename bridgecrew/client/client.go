package bridgecrew

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
	base   string
	token  string
}

