package main

import "github.com/DixonOrtiz/BcryptLoginJson/functions"

/**
 * @DixonOrtiz & @DiegoSepulveda
 * Route only available for admins
 * @description Define if the user accesing this route is an admin
 * @param  {object} req Request
 * @param  {object} res Response
 * @param  {Function} next Callback function
 */

func main() {

	functions.SelectOption()
}
