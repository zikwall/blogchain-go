SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `blogchain`
--

-- --------------------------------------------------------

--
-- Структура таблицы `content`
--

CREATE TABLE `content` (
  `id` int(11) NOT NULL,
  `uuid` varchar(36) NOT NULL,
  `user_id` int(11) NOT NULL,
  `title` varchar(200) NOT NULL,
  `annotation` text NOT NULL,
  `content` mediumtext NOT NULL,
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `content`
--

INSERT INTO `content` (`id`, `uuid`, `user_id`, `title`, `annotation`, `content`, `created_at`, `updated_at`, `image`) VALUES
(10, '56ded262-84bf-11ea-8b77-00d86106f5f8', 2, 'Велотренажер #Самоизоляция или как угомонить ребенка на карантине', 'Весь мир героически борется с «… заразой коронавирусной» (Путин В.В.) Большинство стран закрывают свои границы, своих граждан закрывают на карантин, вводят комендантский час. Вот и Россию не обошла эта гадость стороной.\r\n\r\nВ сложившейся ситуации с пандемией SARS-CoV-2 (COVID-19) все мы с вами сейчас должны находиться на карантине самоизоляции.\r\n\r\nПоэтому вопрос о том, как найти активное развлечение для детей запертых в четырёх стенах стоит как никогда остро. Надо ещё и постараться чтоб эти четыре стены остались, по-возможности, в целости и сохранности.', '<p>За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад. В этой статье я расскажу, как мы разбираем фильмы на структурные элементы и зачем нам это нужно. В конце есть ссылка на репозиторий Github с кодом алгоритмов и примерами. На монтаже кадры нарезают на группы, которые по задумке режиссёра меняют местами и склеивают обратно. Последовательность кадров от одной монтажной склейки до другой в английском языке называют термином shot. К сожалению, русская терминология неудачная, потому что в ней такие группы тоже называются кадрами. Чтобы не запутаться, давайте использовать английский термин. Только введём русскоязычный вариант: За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад. За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад.</p>', 1587382561, 1587656780, '56ded262-84bf-11ea-8b77-00d86106f5f8.png'),
(11, '56ded74e-84bf-11ea-8b77-00d86106f5f8', 2, 'Как мы научились делить видео на сцены с помощью хитрой математики', 'За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад.\r\n\r\nВ этой статье я расскажу, как мы разбираем фильмы на структурные элементы и зачем нам это нужно. В конце есть ссылка на репозиторий Github с кодом алгоритмов и примерами.', 'За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад.\r\n\r\nВ этой статье я расскажу, как мы разбираем фильмы на структурные элементы и зачем нам это нужно. В конце есть ссылка на репозиторий Github с кодом алгоритмов и примерами.\r\n\r\nНа монтаже кадры нарезают на группы, которые по задумке режиссёра меняют местами и склеивают обратно. Последовательность кадров от одной монтажной склейки до другой в английском языке называют термином shot. К сожалению, русская терминология неудачная, потому что в ней такие группы тоже называются кадрами. Чтобы не запутаться, давайте использовать английский термин. Только введём русскоязычный вариант:\r\n\r\nЗа 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад.\r\n\r\nЗа 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад.', 1587382561, NULL, NULL),
(12, '56ded8ef-84bf-11ea-8b77-00d86106f5f8', 11, 'Проверка: насколько «здоровая» у вас поза, когда работаете за компьютером?', 'Ну ладно, домашними рабочими местами похвастались, теперь давайте о действительно важных вещах — о здоровье. Около 8 часов в сутки мы проводим сидя в одной, часто неудобной позе. Некоторые оригиналы — стоя, — но сути это не меняет. Если не уследить, то через несколько лет незаметно испортится осанка, начнутся проблемы с позвоночником, а оттуда головные боли и прочие неприятности. Мы не претендуем на научность, но пообщавшись с несколькими врачами и перечитав кучу релевантных постов тут на Хабре, соорудили короткий тест, который хотя бы в первом приближении покажет, насколько «здоровая» у вас поза для работы.', '<p class=\"ql-align-justify\"><span class=\"ql-font-monospace\">За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, </span></p><p class=\"ql-align-justify\"><span class=\"ql-font-monospace\">которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад. В этой статье я расскажу, как мы разбираем фильмы на структурные элементы и зачем нам это нужно. В конце есть ссылка на репозиторий Github с кодом алгоритмов и примерами. На монтаже кадры нарезают на группы, которые по задумке режиссёра меняют местами и склеивают обратно. Последовательность кадров от одной монтажной склейки до другой в английском языке называют термином shot. К сожалению, русская терминология неудачная, потому что в ней такие группы тоже называются кадрами. Чтобы не запутаться, давайте использовать английский термин. Только введём русскоязычный вариант: За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад. За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад.</span></p>', 1587382561, 1587850395, ''),
(13, '56deda5c-84bf-11ea-8b77-00d86106f5f8', 2, 'First Content', '', 'Content here!', 1587479093, NULL, ''),
(14, '56dedacf-84bf-11ea-8b77-00d86106f5f8', 2, 'Second Content 2', 'Annotation here III', '<p>Content here 2!</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">fn</span> <span class=\"hljs-title\">Hello</span></span>() -&gt; <span class=\"hljs-built_in\">Option</span> {\n   <span class=\"hljs-keyword\">return</span> <span class=\"hljs-literal\">Some</span>(<span class=\"hljs-number\">1</span>)\n}\n\n<span class=\"hljs-function\"><span class=\"hljs-keyword\">fn</span> <span class=\"hljs-title\">Googby</span></span>() -&gt; <span class=\"hljs-built_in\">Option</span> {\n   <span class=\"hljs-keyword\">return</span> <span class=\"hljs-literal\">Some</span>(<span class=\"hljs-number\">0</span>)\n}\n</pre>', 1587481993, 1587644907, '56dedacf-84bf-11ea-8b77-00d86106f5f8.png'),
(15, 'b903ffea-b5b5-45a8-a148-4ad9afe1a34a', 2, 'Second Content 5', '', 'Content here 5!', 1587577375, NULL, ''),
(16, 'a446fe09-92a7-41b0-8073-08b46ed59be1', 2, 'Second Content 5', '', 'Content here 5!', 1587577387, NULL, ''),
(17, 'dda37d06-a12f-4680-859a-6b8667e3b8df', 2, 'Second Content 5', '', 'Content here 5!', 1587577458, NULL, ''),
(18, '1eb00ab9-ecdc-4629-867a-4303053e9eb8', 2, 'Second Content 5', '', 'Content here 5!', 1587577485, NULL, ''),
(19, '41e980d5-824a-4415-9d03-02437bd01ae8', 2, 'Second Content 5', '', 'Content here 5!', 1587577500, NULL, ''),
(20, 'ba0e0428-9bad-4512-a87b-9b9c71b744ce', 2, 'Test I', 'Annotation I', '<pre class=\"ql-syntax\" spellcheck=\"false\">fn Hello() -&gt; Option {\n   return Some(1)\n}\n</pre>', 1587583772, 1587844487, 'ba0e0428-9bad-4512-a87b-9b9c71b744ce.png'),
(21, 'b29ccb38-0b20-43f6-95a8-9bd86cb396b4', 2, 'Test with tags', 'Test with tags', '<h1>Hello</h1><p><br></p><h2>Hi</h2><h2><br></h2><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">func</span> <span class=\"hljs-title\">HiGopher</span>() <span class=\"hljs-title\">HelloResponse</span>\n  <span class=\"hljs-title\">return</span> <span class=\"hljs-title\">Hello</span>\n}\n</span></pre><p><br></p><p><span class=\"ql-formula\" data-value=\"e=mc^2\">﻿<span contenteditable=\"false\"><span class=\"katex\"><span class=\"katex-mathml\"><math xmlns=\"http://www.w3.org/1998/Math/MathML\"><semantics><mrow><mi>e</mi><mo>=</mo><mi>m</mi><msup><mi>c</mi><mn>2</mn></msup></mrow><annotation encoding=\"application/x-tex\">e=mc^2</annotation></semantics></math></span><span class=\"katex-html\" aria-hidden=\"true\"><span class=\"base\"><span class=\"strut\" style=\"height: 0.43056em; vertical-align: 0em;\"></span><span class=\"mord mathdefault\">e</span><span class=\"mspace\" style=\"margin-right: 0.277778em;\"></span><span class=\"mrel\">=</span><span class=\"mspace\" style=\"margin-right: 0.277778em;\"></span></span><span class=\"base\"><span class=\"strut\" style=\"height: 0.814108em; vertical-align: 0em;\"></span><span class=\"mord mathdefault\">m</span><span class=\"mord\"><span class=\"mord mathdefault\">c</span><span class=\"msupsub\"><span class=\"vlist-t\"><span class=\"vlist-r\"><span class=\"vlist\" style=\"height: 0.814108em;\"><span class=\"\" style=\"top: -3.063em; margin-right: 0.05em;\"><span class=\"pstrut\" style=\"height: 2.7em;\"></span><span class=\"sizing reset-size6 size3 mtight\"><span class=\"mord mtight\">2</span></span></span></span></span></span></span></span></span></span></span></span>﻿</span> </p>', 1587845791, NULL, 'b29ccb38-0b20-43f6-95a8-9bd86cb396b4.png'),
(22, 'abb78e71-5760-48a5-86df-2effe61f57b6', 2, 'Test with tags II', 'Test with tags II', '<p>Test with tags II</p>', 1587845965, NULL, 'abb78e71-5760-48a5-86df-2effe61f57b6.png'),
(23, '949014ec-9d50-4215-9728-2c258ec87429', 2, 'Test', 'Test', '<p>Hello</p><p><br></p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title\">hello</span>(<span class=\"hljs-params\">$php</span>) : <span class=\"hljs-title\">HelloResponse</span> </span>{\n   <span class=\"hljs-keyword\">return</span> <span class=\"hljs-keyword\">new</span> HelloResponse($php);\n}\n</pre>', 1589552187, NULL, '949014ec-9d50-4215-9728-2c258ec87429.png'),
(24, 'c49389fd-3d9b-41a8-8a31-5eea74c9b0d8', 2, 'wswssxxxxx', 'wswsw', '<pre class=\"ql-syntax\" spellcheck=\"false\">function bye($php) : ByeResponse {\n   return new ByeResponse ($php);\n}\n\nfn Okk(ok: int32) : OkHttp -&gt; {\n  return Ok?\n}\n\n\n\n</pre><p>Bye!</p>', 1589552505, 1589560896, 'c49389fd-3d9b-41a8-8a31-5eea74c9b0d8.png');

-- --------------------------------------------------------

--
-- Структура таблицы `content_tag`
--

CREATE TABLE `content_tag` (
  `id` int(11) NOT NULL,
  `content_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `content_tag`
--

INSERT INTO `content_tag` (`id`, `content_id`, `tag_id`) VALUES
(1, 10, 1),
(3, 10, 2),
(4, 10, 3),
(5, 14, 2),
(41, 20, 3),
(42, 20, 8),
(43, 20, 12),
(44, 22, 9),
(45, 22, 2),
(46, 22, 13),
(47, 22, 10),
(48, 22, 5),
(49, 22, 7),
(50, 22, 8),
(51, 22, 11),
(52, 23, 5),
(53, 23, 8),
(54, 23, 2),
(97, 24, 9),
(98, 24, 5),
(99, 24, 2);

-- --------------------------------------------------------

--
-- Структура таблицы `migrations`
--

CREATE TABLE `migrations` (
  `id` varchar(255) NOT NULL,
  `applied_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `migrations`
--

INSERT INTO `migrations` (`id`, `applied_at`) VALUES
('20200411204407-user.sql', '2020-04-11 17:50:51'),
('20200411204705-profile.sql', '2020-04-11 17:50:51'),
('20200412154007-fix_user_id_inc.sql', '2020-04-12 12:41:18'),
('20200419132311-profile_fields.sql', '2020-04-19 10:25:55'),
('20200419152626-profile_suto_inc.sql', '2020-04-19 12:30:49'),
('20200419163959-content.sql', '2020-04-19 13:41:52'),
('20200419182416-content_lenght.sql', '2020-04-19 15:25:01'),
('20200420095229-content_annotation.sql', '2020-04-20 06:54:28'),
('20200420095655-content_annotation_type.sql', '2020-04-20 06:57:55'),
('20200420142834-add_fields_to_content.sql', '2020-04-20 11:36:01'),
('20200422202517-uuid_content.sql', '2020-04-22 17:33:12'),
('20200424092157-content_tags.sql', '2020-04-24 06:39:11'),
('20200425212158-tag_label.sql', '2020-04-25 18:24:07');

-- --------------------------------------------------------

--
-- Структура таблицы `profile`
--

CREATE TABLE `profile` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `public_email` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `profile`
--

INSERT INTO `profile` (`id`, `user_id`, `name`, `public_email`, `avatar`, `location`, `status`, `description`) VALUES
(1, 11, 'Andrey Ka', 'and@m.ru', '', '', 'Study...', ''),
(2, 2, 'Andrey Ka', 'andreyka@mail.blog', 'https://avatars1.githubusercontent.com/u/23422968', 'Russia, Moscow', 'Study...', '#PHP, #Go, #JS, #React, #ReactNative - full stack developer TODO: #Rust');

-- --------------------------------------------------------

--
-- Структура таблицы `tags`
--

CREATE TABLE `tags` (
  `id` int(11) NOT NULL,
  `name` varchar(35) NOT NULL,
  `label` varchar(35) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `tags`
--

INSERT INTO `tags` (`id`, `name`, `label`) VALUES
(1, 'Программирование', 'programming'),
(2, 'Алгоритмы', 'algorithms'),
(3, 'Обработка изображений', 'image_processing'),
(5, 'Управление проектами', '\r\nproject_management'),
(6, 'Робототехника', 'robotics'),
(7, 'Управление разработкой', 'development_management'),
(8, 'Системное администрирование', 'devops'),
(9, 'Open source', 'open_source'),
(10, 'Интернет вещей', 'internet_of_things'),
(11, 'Графичекий дизайн', 'design'),
(12, 'Компьютерное железо', 'hardware'),
(13, 'Сетевые технологии', 'network_technologies');

-- --------------------------------------------------------

--
-- Структура таблицы `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(25) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(60) NOT NULL,
  `confirmed_at` int(11) DEFAULT NULL,
  `blocked_at` int(11) DEFAULT NULL,
  `registration_ip` varchar(255) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `user`
--

INSERT INTO `user` (`id`, `username`, `email`, `password_hash`, `confirmed_at`, `blocked_at`, `registration_ip`, `created_at`, `updated_at`) VALUES
(2, 'zikwall', 'dj-on-ik@mail.ru', '$2a$10$46cvIjfMR5vkkISOjhOPbuhy/yPZM5qseNct9ykHCuE3axeUIvA8O', NULL, NULL, '', 1586695520, NULL);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `content`
--
ALTER TABLE `content`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `content_tag`
--
ALTER TABLE `content_tag`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `profile`
--
ALTER TABLE `profile`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `tags`
--
ALTER TABLE `tags`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_unique_username` (`username`),
  ADD UNIQUE KEY `user_unique_email` (`email`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `content`
--
ALTER TABLE `content`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT для таблицы `content_tag`
--
ALTER TABLE `content_tag`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=100;

--
-- AUTO_INCREMENT для таблицы `profile`
--
ALTER TABLE `profile`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT для таблицы `tags`
--
ALTER TABLE `tags`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT для таблицы `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;