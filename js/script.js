const port = 8081
const domain = 'localhost'

function getStatusDescription(status) {
	switch (status) {
		// Informational responses
		case 100:
			return {
				message: 'Continue - Клиент должен продолжить запрос.',
				className: 'special-info',
			}
		case 101:
			return {
				message: 'Switching Protocols - Сервер переключается на другой протокол по запросу клиента.',
				className: 'info',
			}
		case 102:
			return {
				message: 'Processing - Сервер принял запрос, но обработка не завершена.',
				className: 'info',
			}
		case 103:
			return {
				message: 'Early Hints - Сервер готов отправить заголовки ответа до того, как весь запрос будет завершен.',
				className: 'info',
			}
		// Success responses
		case 200:
			return {
				message: 'OK - Запрос выполнен успешно.',
				className: 'success',
			}
		case 201:
			return {
				message: 'Created - Запрос выполнен успешно, и создан новый ресурс.',
				className: 'success',
			}
		case 202:
			return {
				message: 'Accepted - Запрос принят, но обработка ещё не завершена.',
				className: 'success',
			}
		case 203:
			return {
				message: 'Non-Authoritative Information - Ответ содержит измененные данные с сервера-посредника.',
				className: 'success',
			}
		case 204:
			return {
				message: 'No Content - Запрос выполнен успешно, но нет содержимого для возврата.',
				className: 'success',
			}
		case 205:
			return {
				message: 'Reset Content - Запрос выполнен успешно, и клиенту нужно сбросить форму отправки.',
				className: 'success',
			}
		case 206:
			return {
				message:
					'Partial Content - Сервер вернул только часть ресурса из-за заголовка диапазона, отправленного клиентом.',
				className: 'success',
			}
		case 207:
			return {
				message: 'Multi-Status - Тело ответа передает информацию о нескольких статусах для разных операций.',
				className: 'success',
			}
		case 208:
			return {
				message: 'Already Reported - Результаты уже были сообщены в предыдущем ответе на запрос.',
				className: 'success',
			}
		// Redirection messages
		case 300:
			return {
				message: 'Multiple Choices - Запрос может быть выполнен несколькими способами.',
				className: 'special-redirect',
			}
		case 301:
			return {
				message: 'Moved Permanently - Ресурс был перемещен на новый URL.',
				className: 'redirect',
			}
		case 302:
			return {
				message: 'Found - Ресурс временно доступен по другому URL.',
				className: 'redirect',
			}
		case 303:
			return {
				message: 'See Other - Клиент должен использовать другой URI для получения ресурса.',
				className: 'redirect',
			}
		case 304:
			return {
				message: 'Not Modified - Данные не изменились, загружать их повторно не нужно.',
				className: 'redirect',
			}
		case 305:
			return {
				message: 'Use Proxy - Доступ к запрашиваемому ресурсу должен осуществляться через прокси.',
				className: 'redirect',
			}
		case 306:
			return {
				message: 'Reserved - Код ранее использовался, но сейчас зарезервирован и не должен использоваться.',
				className: 'redirect',
			}
		case 307:
			return {
				message: 'Temporary Redirect - Клиент должен повторить запрос по другому URL временно.',
				className: 'redirect',
			}
		case 308:
			return {
				message: 'Permanent Redirect - Клиент должен использовать новый URL для будущих запросов к этому ресурсу.',
				className: 'redirect',
			}
		// Client error responses
		case 400:
			return {
				message: 'Bad Request - Некорректный запрос. Возможно, отсутствуют необходимые параметры.',
				className: 'client-error',
			}
		case 401:
			return {
				message:
					'Unauthorized - Ошибка аутентификации. Либо аутентификация не выполнена, либо у пользователя нет прав доступа.',
				className: 'client-error',
			}
		case 402:
			return {
				message:
					'Payment Required - Требуется оплата для доступа к ресурсу. Этот код в настоящее время редко используется.',
				className: 'client-error',
			}
		case 403:
			return {
				message: 'Forbidden - Доступ запрещен. Аутентификация выполнена, но у пользователя нет прав доступа.',
				className: 'client-error',
			}
		case 404:
			return {
				message: 'Not Found - Запрашиваемый ресурс не найден.',
				className: 'client-error',
			}
		case 405:
			return {
				message: 'Method Not Allowed - Метод запроса не поддерживается для данного ресурса.',
				className: 'client-error',
			}
		case 406:
			return {
				message: 'Not Acceptable - Запрашиваемый ресурс не может быть предоставлен в приемлемом формате.',
				className: 'client-error',
			}
		case 407:
			return {
				message: 'Proxy Authentication Required - Требуется аутентификация через прокси.',
				className: 'client-error',
			}
		case 408:
			return {
				message: 'Request Timeout - Сервер не дождался завершения запроса от клиента.',
				className: 'client-error',
			}
		case 409:
			return {
				message: 'Conflict - Конфликт в запросе, например, конфликт версий данных.',
				className: 'client-error',
			}
		case 410:
			return {
				message: 'Gone - Ресурс был удален и больше недоступен.',
				className: 'client-error',
			}
		case 411:
			return {
				message: 'Length Required - Необходимо указать длину тела запроса для его обработки.',
				className: 'client-error',
			}
		case 412:
			return {
				message: 'Precondition Failed - Предусловия в заголовке запроса не выполнены.',
				className: 'client-error',
			}
		case 413:
			return {
				message: 'Payload Too Large - Тело запроса превышает максимально допустимый размер.',
				className: 'client-error',
			}
		case 414:
			return {
				message: 'URI Too Long - Запрашиваемый URI слишком длинный для обработки сервером.',
				className: 'client-error',
			}
		case 415:
			return {
				message: 'Unsupported Media Type - Тип данных запроса не поддерживается сервером.',
				className: 'client-error',
			}
		case 416:
			return {
				message: 'Range Not Satisfiable - Запрашиваемый диапазон данных не может быть удовлетворен.',
				className: 'client-error',
			}
		case 417:
			return {
				message:
					'Expectation Failed - Ожидания, указанные в заголовке запроса Expect, не могут быть выполнены сервером.',
				className: 'client-error',
			}
		case 418:
			return {
				message:
					"I'm a teapot - Этим кодом обычно обозначается шутка, ссылающаяся на протокол заваривания чая в HTTP.",
				className: 'client-error',
			}
		case 421:
			return {
				message: 'Misdirected Request - Запрос был направлен на сервер, который не может дать ответ.',
				className: 'client-error',
			}
		case 422:
			return {
				message:
					'Unprocessable Entity - Сервер понимает запрос, но не может его обработать из-за семантических ошибок.',
				className: 'client-error',
			}
		case 423:
			return {
				message: 'Locked - Доступ к ресурсу заблокирован.',
				className: 'client-error',
			}
		case 424:
			return {
				message: 'Failed Dependency - Запрос не выполнен из-за сбоя в предыдущем запросе, от которого он зависит.',
				className: 'client-error',
			}
		case 425:
			return {
				message: 'Too Early - Сервер не готов обработать запрос из-за возможных последующих запросов.',
				className: 'client-error',
			}
		case 426:
			return {
				message: 'Upgrade Required - Клиент должен переключиться на другой протокол для завершения запроса.',
				className: 'client-error',
			}
		case 428:
			return {
				message: 'Precondition Required - Сервер требует, чтобы запрос включал предусловие.',
				className: 'client-error',
			}
		case 429:
			return {
				message: 'Too Many Requests - Клиент отправил слишком много запросов за короткий промежуток времени.',
				className: 'client-error',
			}
		case 431:
			return {
				message: 'Request Header Fields Too Large - Поля заголовка запроса слишком велики для обработки сервером.',
				className: 'client-error',
			}
		// Server error responses
		case 500:
			return {
				message: 'Internal Server Error - Внутренняя ошибка сервера.',
				className: 'server-error',
			}
		case 501:
			return {
				message: 'Not Implemented - Сервер не поддерживает функциональность, необходимую для обработки запроса.',
				className: 'server-error',
			}
		case 502:
			return {
				message: 'Bad Gateway - Некорректный ответ от вышестоящего сервера.',
				className: 'server-error',
			}
		case 503:
			return {
				message:
					'Service Unavailable - Сервис временно недоступен. Возможно, сервер перегружен или находится на обслуживании.',
				className: 'server-error',
			}
		case 504:
			return {
				message: 'Gateway Timeout - Сервер не дождался ответа от вышестоящего сервера в установленное время.',
				className: 'server-error',
			}
		case 505:
			return {
				message: 'HTTP Version Not Supported - Сервер не поддерживает указанную в запросе версию протокола HTTP.',
				className: 'server-error',
			}
		case 506:
			return {
				message:
					'Variant Also Negotiates - Ошибка при обработке несколькими вариантами одного ресурса, сервер не может выбрать подходящий.',
				className: 'server-error',
			}
		case 507:
			return {
				message: 'Insufficient Storage - Сервер не может завершить запрос из-за нехватки памяти для обработки.',
				className: 'server-error',
			}
		case 508:
			return {
				message: 'Loop Detected - Сервер обнаружил бесконечный цикл при обработке запроса.',
				className: 'server-error',
			}
		case 510:
			return {
				message: 'Not Extended - Сервер требует дальнейшего расширения запроса для его обработки.',
				className: 'server-error',
			}
		case 511:
			return {
				message: 'Network Authentication Required - Клиент должен пройти аутентификацию в сети для получения доступа.',
				className: 'server-error',
			}
		default:
			return {
				message: 'Unknown Status - Информация о статусе недоступна.',
				className: '',
			}
	}
}

function sendRequest() {
	const url = document.getElementById('urlInput').value
	const method = document.getElementById('methodInput').value
	const data = JSON.parse(document.getElementById('dataInput').value)
	const headers = JSON.parse(document.getElementById('headersInput').value)

	const payload = { url, method, data, headers }

	fetch(`http://${domain}:${port}/send_request/`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(payload),
	})
		.then(response => response.json())
		.then(data => {
			const statusInfo = getStatusDescription(data.status)
			const statusElement = document.getElementById('responseOutput')
			statusElement.textContent = `Status: ${data.status} - ${statusInfo.message}`
			statusElement.className = `status ${statusInfo.className}`

			document.getElementById('responseBody').textContent = JSON.stringify(data.body, null, 2).replace(/\\/g, '')

			addToHistory(payload, data)
		})
		.catch(error => {
			document.getElementById('responseOutput').textContent = ''
			document.getElementById('responseBody').textContent = 'Error sending request'
		})
}

function addToHistory(payload, response) {
	const historyList = document.getElementById('requestHistory')

	const listItem = document.createElement('li')
	listItem.style.position = 'relative'

	listItem.innerHTML = `
        <div style="display: flex; justify-content: space-between; align-items: center;">
            <div>
                <strong>URL:</strong> ${payload.url}<br>
                <strong>Method:</strong> ${payload.method}<br>
                <strong>Status Code:</strong> ${response.status}
            </div>
            <button class="delete-btn" style="background-color: transparent; color: #ff6f61; border: none; cursor: pointer; position: absolute; top: 0; right: 0; padding: 5px; display: none;">✖</button>
        </div>
        <div class="details" style="display: none;">
            <strong>Data:</strong> <pre>${convertToJsonString(payload.data)}</pre><br>
            <strong>Headers:</strong> <pre>${convertToJsonString(payload.headers)}</pre><br>
            <strong>Response:</strong> <pre>${JSON.stringify(response.body, null, 2).replace(/\\/g, '')}</pre>
        </div>
    `

	listItem.addEventListener('click', () => {
		const details = listItem.querySelector('.details')
		details.style.display = details.style.display === 'none' || details.style.display === '' ? 'block' : 'none'
	})

	const deleteBtn = listItem.querySelector('.delete-btn')
	deleteBtn.addEventListener('click', e => {
		e.stopPropagation()
		historyList.removeChild(listItem)
	})

	listItem.addEventListener('mouseenter', () => {
		deleteBtn.style.display = 'inline'
	})
	listItem.addEventListener('mouseleave', () => {
		deleteBtn.style.display = 'none'
	})

	historyList.appendChild(listItem)
}

function convertToJsonString(input) {
	try {
		if (typeof input === 'object' && input !== null) {
			return JSON.stringify(input, null, 2)
		} else {
			return input
		}
	} catch (e) {
		return input
	}
}
