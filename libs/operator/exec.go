package operator

import (
  "github.com/codegangsta/cli"
)

func ExecOperatorInfo(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.OperatorInfo()
  if err != nil {
    return nil, "", err
  }

  return nil, raw, nil
}

func ExecVerify(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.OperatorVerify()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}

func ExecOperators(c *cli.Context) (*Response, string, error){
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.Operators()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}

func ExecSupportToken(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.OperatorSupportToken()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}

func ExecToken(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.OperatorToken()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}

func ExecPassword(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.OperatorPassword()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}
